import 'dart:async';
import 'dart:convert';
import 'package:web_socket_channel/web_socket_channel.dart';

import '../../domain/repositories/websocket_repository.dart';
import '../models/auction_state_model.dart';
import '../models/ws_message_model.dart';

class WebSocketRepositoryImpl implements WebSocketRepository {
  final String baseUrl;
  WebSocketChannel? _channel;
  
  final _messageController = StreamController<dynamic>.broadcast();

  WebSocketRepositoryImpl({required this.baseUrl});

  @override
  Future<void> connect(String url) async {
    _channel = WebSocketChannel.connect(Uri.parse(url));

    _channel!.stream.listen(
      (message) {
        _onMessageReceived(message);
      },
      onError: (err) {
        _messageController.addError(err);
      },
      onDone: () {
      },
    );
  }

  void _onMessageReceived(String rawRespone) {
    try {
      final json = jsonDecode(rawRespone);
      final messageEnvelope = WsMessageModel.fromJson(json);

      switch (messageEnvelope.action) {
        case 'room_state':
          final stateModel = AuctionStateEventModel.fromJson(messageEnvelope.payload);
          _messageController.add(stateModel); 
          break;

        case 'new_bid':
          final bidModel = BidEventModel.fromJson(messageEnvelope.payload, messageEnvelope.roomId);
          _messageController.add(bidModel); 
          break;

        case 'viewer_count':
          final viewerModel = ViewerEventModel.fromJson(messageEnvelope.payload);
          _messageController.add(viewerModel); 
          break;

        default:
          print("Unknown action: ${messageEnvelope.action}");
      }
    } catch (e) {
      print("Parsing WebSocket Error: $e");
    }
  }

  @override
  Stream<dynamic> get messages => _messageController.stream;

  @override
  void send(dynamic data) {
    _channel?.sink.add(jsonEncode(data));
  }

  @override
  Future<void> disconnect() async {
    await _channel?.sink.close();
  }
}