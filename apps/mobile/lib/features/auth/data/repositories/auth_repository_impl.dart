import 'dart:convert';
import 'dart:io';
import 'dart:async';
import 'package:http/http.dart' as http;
import '../../domain/entities/user.dart';
import '../../domain/repositories/auth_repository.dart';
import '../models/user_model.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class AuthRepositoryImpl implements AuthRepository {
  final String baseUrl;
  final http.Client client;
  final FlutterSecureStorage secureStorage;

  AuthRepositoryImpl({
    required this.baseUrl,
    http.Client? client,
    this.secureStorage = const FlutterSecureStorage(),
  }) : client = client ?? http.Client();

  Future<void> _saveTokens(String accessToken, String refreshToken) async {
    await secureStorage.write(key: 'access_token', value: accessToken);
    await secureStorage.write(key: 'refresh_token', value: refreshToken);
  }

  @override
  Future<User?> getCurrentUser() async {
    try {
      final token = await secureStorage.read(key: 'access_token');
      if (token == null) return null;

      final response = await client.get(
        Uri.parse('$baseUrl/auth/me'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      ).timeout(const Duration(seconds: 10));

      if (response.statusCode == 200) {
        final Map<String, dynamic> responseBody = jsonDecode(response.body);
        return UserModel.fromJson(responseBody['data']['user']);
      }

      return null;
    } catch (e) {
      return null;
    }
  }

  @override
  Future<User> login(String email, String password) async {
    try {
      final response = await client.post(
        Uri.parse('$baseUrl/auth/login'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'email': email, 'password': password}),
      ).timeout(const Duration(seconds: 10)); 

      final Map<String, dynamic> responseBody = jsonDecode(response.body);

      if (response.statusCode == 200) {
        final data = responseBody['data'];
        await _saveTokens(data['access_token'], data['refresh_token']);
        return UserModel.fromJson(data['user']);
      } else {
        throw responseBody['message'] ?? 'Login failed';
      }
    } on SocketException {
      throw 'Unable to connect to the server';
    } on TimeoutException {
      throw 'Connection timeout';
    } catch (e) {
      throw e.toString();
    }
  }

  @override
  Future<void> register(String email, String password, String fullName) async {
    try {
    final response = await client.post(
      Uri.parse('$baseUrl/auth/register'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'email': email,
        'password': password,
        'full_name': fullName,
      }),
    ).timeout(const Duration(seconds: 10))  ;

    if (response.statusCode == 201 || response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return;
    } else {
      final error = jsonDecode(response.body);
      throw Exception(error['message'] ?? 'Register failed');
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