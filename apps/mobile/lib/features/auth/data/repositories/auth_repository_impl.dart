import 'dart:convert';
import 'package:http/http.dart' as http;
import '../../domain/entities/user.dart';
import '../../domain/repositories/auth_repository.dart';
import '../models/user_model.dart';
import '../../core/utils/env_config.dart';

class AuthRepositoryImpl implements AuthRepository {
  final String baseUrl;
  final http.Client client;

  AuthRepositoryImpl({required this.baseUrl, http.Client? client}) 
      : client = client ?? http.Client();

  @override
  Future<User?> getCurrentUser() async {
    return null; 
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
        return UserModel.fromJson(responseBody['data']['user']);
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