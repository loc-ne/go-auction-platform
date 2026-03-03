import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'features/auth/data/repositories/auth_repository_impl.dart';
import 'features/auth/domain/repositories/auth_repository.dart';


void main() {
  const String apiUrl = "http://10.0.2.2:8080/api/v1";
  final AuthRepository authRepository = AuthRepositoryImpl(baseUrl: apiUrl); 

  runApp(MyApp(authRepository: authRepository));
}


