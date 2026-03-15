import 'dart:convert';
import 'dart:io';
import 'dart:async';
import 'package:http/http.dart' as http;
import '../../domain/repositories/auction_repository.dart';

class AuctionRepositoryImpl implements AuctionRepository {
  final String baseUrl;
  final http.Client client;

  AuctionRepositoryImpl({
    required this.baseUrl,
    http.Client? client,
  }) : client = client ?? http.Client();


  @override
  Future<void> bid(String productId, int amount) async {
    try {
    final response = await client.post(
      Uri.parse('$baseUrl/bids'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'product_id': productId,
        'amount': amount,
      }),
    ).timeout(const Duration(seconds: 10))  ;

    if (response.statusCode == 201 || response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return;
    } else {
      final error = jsonDecode(response.body);
      throw Exception(error['error'] ?? 'Bid failed');
    }
  } on SocketException {
    throw 'Unable to connect to the server';
  } on TimeoutException {
    throw 'Connection timeout';
  } catch (e) {
    throw e.toString();
  }
}

}