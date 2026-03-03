import '../../domain/entities/user.dart';

class UserModel extends User {
  UserModel({required super.email, required super.fullName, required super.role});

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      email: json['email'],
      fullName: json['fullName'],
      role: json['role'],
    );
  }
}
