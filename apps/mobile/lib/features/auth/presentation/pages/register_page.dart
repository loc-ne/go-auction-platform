import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:lucide_icons/lucide_icons.dart';
import '../bloc/auth_bloc.dart';
import '../bloc/auth_event.dart';
import '../bloc/auth_state.dart';

class RegisterPage extends StatefulWidget {
  const RegisterPage({super.key});

  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final _nameController = TextEditingController();
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();

  @override
  void dispose() {
    _nameController.dispose();
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        leading: IconButton(
          icon: const Icon(LucideIcons.arrowLeft, color: Colors.black87),
          onPressed: () => Navigator.pop(context),
        ),
      ),
      body: BlocConsumer<AuthBloc, AuthState>(
        listener: (context, state) {
          if (state is AuthRegisterSuccess) {
            Navigator.pushReplacementNamed(context, '/login');
          }
          if (state is AuthError) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text(state.message)),
            );
          }
        },
        builder: (context, state) {
          return SingleChildScrollView(
            padding: const EdgeInsets.symmetric(horizontal: 24.0, vertical: 16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'Tạo tài khoản\nmới ✨',
                  style: GoogleFonts.outfit(
                    fontSize: 32,
                    fontWeight: FontWeight.w800,
                    color: const Color(0xFF0F172A),
                    height: 1.2,
                  ),
                ),
                const SizedBox(height: 8),
                Text(
                  'Gia nhập cộng đồng Snap & Bid để săn những món đồ giá trị.',
                  style: GoogleFonts.inter(
                    fontSize: 14,
                    color: Colors.grey.shade600,
                  ),
                ),
                const SizedBox(height: 36),

                // Name Input
                Text(
                  'Họ và tên',
                  style: GoogleFonts.inter(fontSize: 13, fontWeight: FontWeight.w600, color: Colors.black87),
                ),
                const SizedBox(height: 8),
                TextField(
                  controller: _nameController,
                  decoration: InputDecoration(
                    hintText: 'Nhập họ và tên',
                    hintStyle: GoogleFonts.inter(color: Colors.grey.shade400, fontSize: 14),
                    filled: true,
                    fillColor: Colors.grey.shade50,
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(12),
                      borderSide: BorderSide(color: Colors.grey.shade200),
                    ),
                    enabledBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(12),
                      borderSide: BorderSide(color: Colors.grey.shade200),
                    ),
                    focusedBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(12),
                      borderSide: const BorderSide(color: Color(0xFF0F172A)),
                    ),
                  ),
                ),
                const SizedBox(height: 16),

                // Email Input
                Text(
                  'Email',
                  style: GoogleFonts.inter(fontSize: 13, fontWeight: FontWeight.w600, color: Colors.black87),
                ),
                const SizedBox(height: 8),
                TextField(
                  controller: _emailController,
                  decoration: InputDecoration(
                    hintText: 'Nhập email của bạn',
                    hintStyle: GoogleFonts.inter(color: Colors.grey.shade400, fontSize: 14),
                    filled: true,
                    fillColor: Colors.grey.shade50,
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(12),
                      borderSide: BorderSide(color: Colors.grey.shade200),
                    ),
                    enabledBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(12),
                      borderSide: BorderSide(color: Colors.grey.shade200),
                    ),
                    focusedBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(12),
                      borderSide: const BorderSide(color: Color(0xFF0F172A)),
                    ),
                  ),
                  keyboardType: TextInputType.emailAddress,
                ),
                const SizedBox(height: 16),

                // Password Input
                Text(
                  'Mật khẩu',
                  style: GoogleFonts.inter(fontSize: 13, fontWeight: FontWeight.w600, color: Colors.black87),
                ),
                const SizedBox(height: 8),
                TextField(
                  controller: _passwordController,
                  decoration: InputDecoration(
                    hintText: 'Nhập mật khẩu',
                    hintStyle: GoogleFonts.inter(color: Colors.grey.shade400, fontSize: 14),
                    filled: true,
                    fillColor: Colors.grey.shade50,
                    border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(12),
                      borderSide: BorderSide(color: Colors.grey.shade200),
                    ),
                    enabledBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(12),
                      borderSide: BorderSide(color: Colors.grey.shade200),
                    ),
                    focusedBorder: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(12),
                      borderSide: const BorderSide(color: Color(0xFF0F172A)),
                    ),
                  ),
                  obscureText: true,
                ),
                const SizedBox(height: 32),

                // Register Button
                SizedBox(
                  width: double.infinity,
                  height: 52,
                  child: ElevatedButton(
                    onPressed: state is AuthLoading
                        ? null
                        : () {
                            context.read<AuthBloc>().add(
                                  AuthRegisterRequested(
                                    email: _emailController.text,
                                    password: _passwordController.text,
                                    name: _nameController.text,
                                  ),
                                );
                          },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: const Color(0xFF0F172A),
                      foregroundColor: Colors.white,
                      elevation: 0,
                      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
                    ),
                    child: state is AuthLoading
                        ? const SizedBox(height: 24, width: 24, child: CircularProgressIndicator(color: Colors.white, strokeWidth: 2))
                        : Text(
                            'ĐĂNG KÝ',
                            style: GoogleFonts.inter(fontWeight: FontWeight.bold, fontSize: 14, letterSpacing: 1.0),
                          ),
                  ),
                ),
                const SizedBox(height: 20),

                // Divider
                Row(
                  children: [
                    Expanded(child: Divider(color: Colors.grey.shade300)),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 16),
                      child: Text('hoặc', style: GoogleFonts.inter(color: Colors.grey.shade500, fontSize: 12)),
                    ),
                    Expanded(child: Divider(color: Colors.grey.shade300)),
                  ],
                ),
                const SizedBox(height: 20),

                // Google Button
                SizedBox(
                  width: double.infinity,
                  height: 52,
                  child: OutlinedButton.icon(
                    onPressed: () {},
                    icon: Image.network(
                      'https://cdn1.iconfinder.com/data/icons/google-s-logo/150/Google_Icons-09-512.png',
                      height: 24,
                    ),
                    label: Text(
                      'Đăng ký bằng Google',
                      style: GoogleFonts.inter(fontWeight: FontWeight.w600, color: Colors.black87),
                    ),
                    style: OutlinedButton.styleFrom(
                      side: BorderSide(color: Colors.grey.shade300),
                      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
                    ),
                  ),
                ),
                const SizedBox(height: 32),

                // Login Redirect
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text('Đã có tài khoản? ', style: GoogleFonts.inter(color: Colors.grey.shade600)),
                    GestureDetector(
                      onTap: () => Navigator.pushReplacementNamed(context, '/login'),
                      child: Text(
                        'Đăng nhập ngay',
                        style: GoogleFonts.inter(fontWeight: FontWeight.bold, color: const Color(0xFF0F172A)),
                      ),
                    ),
                  ],
                ),
                const SizedBox(height: 24),
              ],
            ),
          );
        },
      ),
    );
  }
}
