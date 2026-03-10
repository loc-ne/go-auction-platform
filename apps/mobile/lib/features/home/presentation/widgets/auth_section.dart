import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../../auth/presentation/bloc/auth_bloc.dart';
import '../../../auth/presentation/bloc/auth_state.dart';
import '../../../auth/presentation/bloc/auth_event.dart';
import '../../../auth/presentation/pages/login_page.dart';

class AuthSection extends StatelessWidget {
  const AuthSection({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<AuthBloc, AuthState>(
      builder: (context, state) {
        if (state is AuthAuthenticated) {
          return Container(
            margin: const EdgeInsets.symmetric(horizontal: 20, vertical: 16),
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
            decoration: BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.circular(16),
              border: Border.all(color: Colors.grey.shade200),
            ),
            child: Row(
              children: [
                CircleAvatar(
                  radius: 20,
                  backgroundColor: const Color(0xFF0F172A),
                  child: Text(
                    state.user.fullName.isNotEmpty ? state.user.fullName[0].toUpperCase() : 'U',
                    style: GoogleFonts.inter(
                      fontSize: 16,
                      fontWeight: FontWeight.bold,
                      color: Colors.white,
                    ),
                  ),
                ),
                const SizedBox(width: 12),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        'Xin chào,',
                        style: GoogleFonts.inter(
                          fontSize: 12,
                          color: Colors.grey.shade600,
                        ),
                      ),
                      Text(
                        state.user.fullName,
                        style: GoogleFonts.outfit(
                          fontSize: 16,
                          fontWeight: FontWeight.bold,
                          color: const Color(0xFF0F172A),
                        ),
                        maxLines: 1,
                        overflow: TextOverflow.ellipsis,
                      ),
                    ],
                  ),
                ),
                IconButton(
                  onPressed: () {
                    context.read<AuthBloc>().add(AuthLogoutRequested());
                  },
                  icon: const Icon(LucideIcons.logOut, color: Colors.black87, size: 20),
                ),
              ],
            ),
          );
        } else if (state is AuthUnauthenticated || state is AuthInitial) {
          return Container(
            margin: const EdgeInsets.symmetric(horizontal: 20, vertical: 16),
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 16),
            decoration: BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.circular(16),
              border: Border.all(color: Colors.grey.shade200),
            ),
            child: Row(
              children: [
                Expanded(
                  child: Text(
                    'Đăng nhập để tham gia đấu giá',
                    style: GoogleFonts.inter(
                      fontSize: 14,
                      fontWeight: FontWeight.w600,
                      color: const Color(0xFF0F172A),
                    ),
                  ),
                ),
                const SizedBox(width: 12),
                ElevatedButton(
                  onPressed: () {
                    Navigator.push(
                      context,
                      MaterialPageRoute(builder: (context) => const LoginPage()),
                    );
                  },
                  style: ElevatedButton.styleFrom(
                    backgroundColor: const Color(0xFF0F172A),
                    foregroundColor: Colors.white,
                    elevation: 0,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(12),
                    ),
                  ),
                  child: Text(
                    'Tham gia',
                    style: GoogleFonts.inter(fontWeight: FontWeight.w600, fontSize: 13),
                  ),
                ),
              ],
            ),
          );
        }
        return const SizedBox(
          height: 80, 
          child: Center(child: CircularProgressIndicator(strokeWidth: 2)),
        );
      },
    );
  }
}
