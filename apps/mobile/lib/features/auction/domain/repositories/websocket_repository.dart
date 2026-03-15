abstract class WebSocketRepository {
  Future<void> connect(String url);

  Stream<dynamic> get messages;

  void send(dynamic data);

  Future<void> disconnect();
}