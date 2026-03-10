import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:lucide_icons/lucide_icons.dart';
import '../bloc/auth_bloc.dart';
import '../bloc/auth_event.dart';
import '../bloc/auth_state.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();

  @override
  void dispose() {
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
          if (state is AuthAuthenticated) {
            ScaffoldMessenger.of(context).showSnackBar(
              const SnackBar(content: Text('Đăng nhập thành công!')),
            );
            Navigator.pop(context);
          } else if (state is AuthError) {
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
                  'Chào mừng\ntrở lại 👋',
                  style: GoogleFonts.outfit(
                    fontSize: 32,
                    fontWeight: FontWeight.w800,
                    color: const Color(0xFF0F172A),
                    height: 1.2,
                  ),
                ),
                const SizedBox(height: 8),
                Text(
                  'Đăng nhập để tham gia đấu giá',
                  style: GoogleFonts.inter(
                    fontSize: 14,
                    color: Colors.grey.shade600,
                  ),
                ),
                const SizedBox(height: 48),

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
                const SizedBox(height: 20),

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
                const SizedBox(height: 12),
                Align(
                  alignment: Alignment.centerRight,
                  child: TextButton(
                    onPressed: () {},
                    child: Text(
                      'Quên mật khẩu?',
                      style: GoogleFonts.inter(fontSize: 13, fontWeight: FontWeight.w600, color: const Color(0xFF0F172A)),
                    ),
                  ),
                ),
                const SizedBox(height: 24),

                // Login Button
                SizedBox(
                  width: double.infinity,
                  height: 52,
                  child: ElevatedButton(
                    onPressed: state is AuthLoading
                        ? null
                        : () {
                            context.read<AuthBloc>().add(
                                  AuthLoginRequested(
                                    email: _emailController.text,
                                    password: _passwordController.text,
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
                            'ĐĂNG NHẬP',
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

                // Google Logic
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
                      'Đăng nhập với Google',
                      style: GoogleFonts.inter(fontWeight: FontWeight.w600, color: Colors.black87),
                    ),
                    style: OutlinedButton.styleFrom(
                      side: BorderSide(color: Colors.grey.shade300),
                      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
                    ),
                  ),
                ),
                const SizedBox(height: 32),

                // Register Redirect
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text('Chưa có tài khoản? ', style: GoogleFonts.inter(color: Colors.grey.shade600)),
                    GestureDetector(
                      onTap: () => Navigator.pushReplacementNamed(context, '/register'),
                      child: Text(
                        'Đăng ký ngay',
                        style: GoogleFonts.inter(fontWeight: FontWeight.bold, color: const Color(0xFF0F172A)),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
