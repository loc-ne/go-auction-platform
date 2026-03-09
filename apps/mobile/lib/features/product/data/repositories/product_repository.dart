import 'dart:convert';
import 'dart:io';
import 'dart:async';
import 'package:http/http.dart' as http;
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class ProductRepository {
  final String baseUrl;
  final http.Client client;
  final FlutterSecureStorage secureStorage;

  ProductRepository({
    required this.baseUrl,
    http.Client? client,
    this.secureStorage = const FlutterSecureStorage(),
  }) : client = client ?? http.Client();

  Future<List<String>> submitImages(List<File> images) async {
    try {
      final token = await secureStorage.read(key: 'access_token');
      if (token == null) {
        throw 'Vui lòng đăng nhập để thực hiện chức năng này.';
      }

      var request = http.MultipartRequest('POST', Uri.parse('$baseUrl/media/upload'));
      request.headers['Authorization'] = 'Bearer $token';

      for (var image in images) {
        request.files.add(await http.MultipartFile.fromPath(
          'files', 
          image.path,
        ));
      }

      final streamedResponse = await request.send().timeout(const Duration(seconds: 30));
      final response = await http.Response.fromStream(streamedResponse);

      if (response.statusCode == 200 || response.statusCode == 201) {
        final data = jsonDecode(response.body);
        return List<String>.from(data['data']);
      } else {
        final error = jsonDecode(response.body);
        throw error['message'] ?? error['error'] ?? 'Lỗi khi tải ảnh lên server';
      }
    } on SocketException {
      throw 'Không thể kết nối đến máy chủ';
    } on TimeoutException {
      throw 'Kết nối bị quá hạn';
    } catch (e) {
      if (e.toString() == 'Vui lòng đăng nhập để thực hiện chức năng này.') rethrow;
      throw 'Đã có lỗi xảy ra khi tải ảnh lên: $e';
    }
  }

  Future<void> submitProduct({
    required String name,
    required String description,
    required int startingPrice,
    required int bidIncrement,
    required List<String> images,
    required int durationDays,
  }) async {
    try {
      final token = await secureStorage.read(key: 'access_token');
      if (token == null) {
        throw 'Vui lòng đăng nhập để thực hiện chức năng này.';
      }

      final response = await client.post(
        Uri.parse('$baseUrl/products'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: jsonEncode({
          'name': name,
          'description': description,
          'starting_price': startingPrice,
          'bid_increment': bidIncrement,
          'image_urls': images,
          'start_at': DateTime.now().toIso8601String(),
          'end_at': DateTime.now().add(Duration(days: durationDays)).toIso8601String(),
        }),
      ).timeout(const Duration(seconds: 15));

      if (response.statusCode == 200 || response.statusCode == 201) {
        return;
      } else {
        final error = jsonDecode(response.body);
        throw error['message'] ?? error['error'] ?? 'Lỗi khi đăng sản phẩm';
      }
    } on SocketException {
      throw 'Không thể kết nối đến máy chủ';
    } on TimeoutException {
      throw 'Kết nối bị quá hạn';
    } catch (e) {
      throw e.toString();
    }
  }
}
