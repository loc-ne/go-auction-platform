import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:intl/intl.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../widgets/product_card.dart';
import '../widgets/custom_bottom_nav.dart';
import '../widgets/auth_section.dart';
import '../../../product/presentation/pages/submit_product_page.dart';

final List<Map<String, dynamic>> mockTrending = [
  {
    "id": "uuid-1",
    "name": "Rolex Daytona Panda 116500LN",
    "description": "Tinh hoa đồng hồ cơ học Thụy Sỹ. Tình trạng xuất sắc 99%, hộp sổ thẻ.",
    "current_price": 850000000,
    "image_url": "https://images.unsplash.com/photo-1523170335258-f5ed11844a49?auto=format&fit=crop&q=80&w=800",
    "end_at": DateTime.now().add(const Duration(hours: 2)),
    "stats": {"view_count": "24.5k", "favorite_count": "5.1k", "bid_count": 342},
    "rank": 1
  },
  {
    "id": "uuid-2",
    "name": "Air Jordan 1 Chicago 2015",
    "description": "Huyền thoại streetwear. Size 42, Fullbox, deadstock.",
    "current_price": 25000000,
    "image_url": "https://images.unsplash.com/photo-1600185365483-26d7a4cc7519?auto=format&fit=crop&q=80&w=800",
    "end_at": DateTime.now().add(const Duration(hours: 5)),
    "stats": {"view_count": "12.4k", "favorite_count": "3.2k", "bid_count": 128},
    "rank": 2
  },
  {
    "id": "uuid-3",
    "name": "Leica M6 TTL Black",
    "description": "Máy ảnh film Leica huyền thoại. Kèm ống kính 35mm.",
    "current_price": 120000000,
    "image_url": "https://images.unsplash.com/photo-1516961642265-531546e84af2?auto=format&fit=crop&q=80&w=800",
    "end_at": DateTime.now().add(const Duration(hours: 12)),
    "stats": {"view_count": "8.9k", "favorite_count": "2.1k", "bid_count": 85},
    "rank": 3
  },
];

String formatCurrency(int amount) {
  final format = NumberFormat.currency(locale: 'vi_VN', symbol: '₫');
  return format.format(amount);
}

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _currentIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFFF8FAFC), 
      body: CustomScrollView(
        slivers: [
          _buildAppBar(),
          SliverToBoxAdapter(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const AuthSection(),
                _buildHeroBanner(),
                const SizedBox(height: 32),
                _buildTrendingSection(),
                const SizedBox(height: 100), 
              ],
            ),
          ),
        ],
      ),
      extendBody: true, 
      bottomNavigationBar: CustomBottomNav(
        currentIndex: _currentIndex,
        onTap: (index) => setState(() => _currentIndex = index),
        onFabTap: () {
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => const SubmitProductPage(),
            ),
          );
        },
      ),
    );
  }

  Widget _buildAppBar() {
    return SliverAppBar(
      backgroundColor: Colors.white.withOpacity(0.8),
      pinned: true,
      elevation: 0,
      centerTitle: true,
      title: Text(
        'SNAP & BID',
        style: GoogleFonts.outfit(
          color: const Color(0xFF0F172A), 
          fontWeight: FontWeight.w900,
          fontSize: 20,
          letterSpacing: 2.0,
        ),
      ),
      leading: IconButton(
        icon: const Icon(LucideIcons.search, color: Color(0xFF0F172A), size: 22),
        onPressed: () {},
      ),
      actions: [
        IconButton(
          icon: const Icon(LucideIcons.bell, color: Color(0xFF0F172A), size: 22),
          onPressed: () {},
        ),
        const SizedBox(width: 8),
      ],
    );
  }

  Widget _buildHeroBanner() {
    final top1 = mockTrending.first;

    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20.0, vertical: 12.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              Container(
                padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 6),
                decoration: BoxDecoration(
                  color: Colors.orange.shade50,
                  border: Border.all(color: Colors.orange.shade200),
                  borderRadius: BorderRadius.circular(20),
                ),
                child: Row(
                  children: [
                    Icon(LucideIcons.trendingUp, size: 14, color: Colors.orange.shade700),
                    const SizedBox(width: 4),
                    Text(
                      'GIÁ ĐANG SỐT',
                      style: GoogleFonts.inter(
                        fontSize: 10,
                        fontWeight: FontWeight.bold,
                        color: Colors.orange.shade700,
                        letterSpacing: 1.0,
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
          const SizedBox(height: 16),

          Stack(
            children: [
              Container(
                height: 320,
                width: double.infinity,
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(24),
                  border: Border.all(color: Colors.grey.shade200),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.grey.shade200.withOpacity(0.5),
                      blurRadius: 20,
                      offset: const Offset(0, 10),
                    ),
                  ],
                  image: DecorationImage(
                    image: NetworkImage(top1['image_url']),
                    fit: BoxFit.cover,
                  ),
                ),
              ),
              Positioned(
                top: 16,
                left: 16,
                child: Container(
                  padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
                  decoration: BoxDecoration(
                    color: Colors.white.withOpacity(0.9),
                    borderRadius: BorderRadius.circular(16),
                    boxShadow: [
                      BoxShadow(
                        color: Colors.black.withOpacity(0.05),
                        blurRadius: 10,
                      ),
                    ],
                  ),
                  child: Row(
                    children: [
                      const Icon(LucideIcons.crown, size: 14, color: Colors.amber),
                      const SizedBox(width: 4),
                      Text(
                        'RANK 01',
                        style: GoogleFonts.inter(
                          fontSize: 10,
                          fontWeight: FontWeight.bold,
                          color: Colors.black87,
                          letterSpacing: 1.0,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              Positioned(
                bottom: 16,
                right: 16,
                child: Container(
                  padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
                  decoration: BoxDecoration(
                    color: Colors.white.withOpacity(0.95),
                    borderRadius: BorderRadius.circular(12),
                  ),
                  child: Row(
                    children: [
                      const Icon(LucideIcons.clock, size: 14, color: Colors.deepOrange),
                      const SizedBox(width: 6),
                      Text(
                        '02:15:30',
                        style: GoogleFonts.jetBrainsMono(
                          fontSize: 12,
                          fontWeight: FontWeight.bold,
                          color: Colors.black87,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ),
          const SizedBox(height: 20),

          // Tên SP
          Text(
            top1['name'],
            style: GoogleFonts.inter(
              fontSize: 24,
              fontWeight: FontWeight.w900,
              color: const Color(0xFF0F172A),
              height: 1.2,
            ),
          ),
          const SizedBox(height: 8),
          
          // Row Stats
          Row(
            children: [
              Icon(LucideIcons.eye, size: 14, color: Colors.grey.shade500),
              const SizedBox(width: 4),
              Text(top1['stats']['view_count'], style: GoogleFonts.inter(fontSize: 12, color: Colors.grey.shade500, fontWeight: FontWeight.w600)),
              const SizedBox(width: 16),
              Icon(LucideIcons.heart, size: 14, color: Colors.red.shade400),
              const SizedBox(width: 4),
              Text(top1['stats']['favorite_count'], style: GoogleFonts.inter(fontSize: 12, color: Colors.grey.shade500, fontWeight: FontWeight.w600)),
            ],
          ),
          const SizedBox(height: 20),

          // Giá & Nút Bấm
          Row(
            crossAxisAlignment: CrossAxisAlignment.end,
            children: [
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'GIÁ HIỆN TẠI',
                      style: GoogleFonts.inter(fontSize: 10, fontWeight: FontWeight.bold, color: Colors.grey.shade400, letterSpacing: 1.0),
                    ),
                    const SizedBox(height: 4),
                    Text(
                      formatCurrency(top1['current_price']),
                      style: GoogleFonts.jetBrainsMono(
                        fontSize: 24,
                        fontWeight: FontWeight.w800,
                        color: Colors.green.shade700,
                        letterSpacing: -0.5,
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
          const SizedBox(height: 16),
          
          // Nút Đấu Giá vuốt / bấm nảy
          SizedBox(
            width: double.infinity,
            height: 56,
            child: ElevatedButton(
              style: ElevatedButton.styleFrom(
                backgroundColor: const Color(0xFF0F172A),
                foregroundColor: Colors.white,
                shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
                elevation: 0,
              ),
              onPressed: () {
                // TODO: Hiển thị BottomSheet vuốt để chốt giá
              },
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Icon(LucideIcons.gavel, size: 20),
                  const SizedBox(width: 8),
                  Text(
                    'THẢ GIÁ NGAY',
                    style: GoogleFonts.inter(
                      fontSize: 14,
                      fontWeight: FontWeight.bold,
                      letterSpacing: 1.5,
                    ),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildTrendingSection() {
    final list = mockTrending.skip(1).toList();

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20.0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.end,
            children: [
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    'BẢNG XẾP HẠNG',
                    style: GoogleFonts.outfit(
                      fontSize: 18,
                      fontWeight: FontWeight.w900,
                      color: const Color(0xFF0F172A),
                      letterSpacing: 1.0,
                    ),
                  ),
                  const SizedBox(height: 4),
                  Text(
                    'Các bảo vật đang được săn lùng',
                    style: GoogleFonts.inter(
                      fontSize: 12,
                      fontWeight: FontWeight.w500,
                      color: Colors.grey.shade500,
                    ),
                  ),
                ],
              ),
              Text(
                'Xem tất cả',
                style: GoogleFonts.inter(
                  fontSize: 12,
                  fontWeight: FontWeight.bold,
                  color: Colors.blue.shade600,
                ),
              ),
            ],
          ),
        ),
        const SizedBox(height: 16),
        SizedBox(
          height: 380, 
          child: ListView.builder(
            scrollDirection: Axis.horizontal,
            padding: const EdgeInsets.symmetric(horizontal: 16),
            itemCount: list.length,
            itemBuilder: (context, index) {
              final item = list[index];
              return ProductCard(item: item);
            },
          ),
        ),
      ],
    );
  }
}
