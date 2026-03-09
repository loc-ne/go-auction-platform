import 'dart:io';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:lucide_icons/lucide_icons.dart';
import 'package:image_picker/image_picker.dart';

class SubmitProductPage extends StatefulWidget {
  const SubmitProductPage({super.key});

  @override
  State<SubmitProductPage> createState() => _SubmitProductPageState();
}

class _SubmitProductPageState extends State<SubmitProductPage> {
  final _formKey = GlobalKey<FormState>();
  final ImagePicker _picker = ImagePicker();
  List<File> _images = [];

  // Controllers
  final _nameController = TextEditingController();
  final _descController = TextEditingController();
  final _startPriceController = TextEditingController();
  final _bidIncrementController = TextEditingController();

  Future<void> _pickImages() async {
    final List<XFile> pickedFiles = await _picker.pickMultiImage();
    if (pickedFiles.isNotEmpty) {
      setState(() {
        _images.addAll(pickedFiles.map((x) => File(x.path)));
      });
    }
  }

  void _removeImage(int index) {
    setState(() {
      _images.removeAt(index);
    });
  }

  @override
  void dispose() {
    _nameController.dispose();
    _descController.dispose();
    _startPriceController.dispose();
    _bidIncrementController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        title: Text(
          'ĐĂNG SẢN PHẨM',
          style: GoogleFonts.outfit(
            fontWeight: FontWeight.w900,
            fontSize: 18,
            letterSpacing: 1.5,
            color: const Color(0xFF0F172A),
          ),
        ),
        backgroundColor: Colors.white,
        elevation: 0,
        centerTitle: true,
        leading: IconButton(
          icon: const Icon(LucideIcons.x, color: Color(0xFF0F172A), size: 24),
          onPressed: () => Navigator.pop(context),
        ),
      ),
      body: SafeArea(
        child: Form(
          key: _formKey,
          child: ListView(
            padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
            children: [
              // 1. Upload ảnh (Media)
              Text(
                'Hình ảnh Sản phẩm',
                style: GoogleFonts.inter(fontSize: 14, fontWeight: FontWeight.bold, color: const Color(0xFF0F172A)),
              ),
              const SizedBox(height: 8),
              Text(
                'Tải lên ít nhất 1 hình ảnh rõ nét của sản phẩm (tối đa 5 ảnh).',
                style: GoogleFonts.inter(fontSize: 12, color: Colors.grey.shade500),
              ),
              const SizedBox(height: 16),
              
              SizedBox(
                height: 100,
                child: ListView(
                  scrollDirection: Axis.horizontal,
                  children: [
                    // Nút thêm ảnh
                    GestureDetector(
                      onTap: _pickImages,
                      child: Container(
                        width: 100,
                        margin: const EdgeInsets.only(right: 12),
                        decoration: BoxDecoration(
                          color: Colors.grey.shade50,
                          border: Border.all(color: Colors.grey.shade300, style: BorderStyle.solid),
                          borderRadius: BorderRadius.circular(16),
                        ),
                        child: Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            Icon(LucideIcons.camera, color: Colors.grey.shade400, size: 28),
                            const SizedBox(height: 8),
                            Text('Thêm ảnh', style: GoogleFonts.inter(fontSize: 12, color: Colors.grey.shade500)),
                          ],
                        ),
                      ),
                    ),
                    // Danh sách ảnh đã chọn
                    ...List.generate(_images.length, (index) {
                      return Stack(
                        children: [
                          Container(
                            width: 100,
                            margin: const EdgeInsets.only(right: 12),
                            decoration: BoxDecoration(
                              borderRadius: BorderRadius.circular(16),
                              image: DecorationImage(
                                image: FileImage(_images[index]),
                                fit: BoxFit.cover,
                              ),
                            ),
                          ),
                          Positioned(
                            top: 4,
                            right: 16,
                            child: GestureDetector(
                              onTap: () => _removeImage(index),
                              child: Container(
                                padding: const EdgeInsets.all(4),
                                decoration: const BoxDecoration(
                                  color: Colors.black54,
                                  shape: BoxShape.circle,
                                ),
                                child: const Icon(LucideIcons.x, size: 14, color: Colors.white),
                              ),
                            ),
                          ),
                        ],
                      );
                    }),
                  ],
                ),
              ),
              
              const SizedBox(height: 32),

              // 2. Tên Sản phẩm
              _buildInputLabel('Tên Sản Phẩm'),
              TextFormField(
                controller: _nameController,
                decoration: _buildInputDecoration('Nhập tên sản phẩm (VD: Rolex Daytona panda...)'),
                validator: (val) => val == null || val.isEmpty ? 'Vui lòng nhập tên sản phẩm' : null,
              ),
              
              const SizedBox(height: 24),

              // 3. Mô tả
              _buildInputLabel('Mô tả Chi Tiết'),
              TextFormField(
                controller: _descController,
                maxLines: 4,
                decoration: _buildInputDecoration('Mô tả tình trạng, năm sản xuất, phụ kiện đi kèm...'),
                validator: (val) => val == null || val.trim().isEmpty ? 'Vui lòng nhập mô tả' : null,
              ),

              const SizedBox(height: 24),

              // 4. Giá & Bước Giá (Row)
              Row(
                children: [
                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        _buildInputLabel('Giá Khởi Điểm (VNĐ)'),
                        TextFormField(
                          controller: _startPriceController,
                          keyboardType: TextInputType.number,
                          decoration: _buildInputDecoration('0 ₫'),
                          validator: (val) => val == null || val.isEmpty ? 'Bắt buộc' : null,
                        ),
                      ],
                    ),
                  ),
                  const SizedBox(width: 16),
                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        _buildInputLabel('Bước Giá (VNĐ)'),
                        TextFormField(
                          controller: _bidIncrementController,
                          keyboardType: TextInputType.number,
                          decoration: _buildInputDecoration('0 ₫'),
                          validator: (val) => val == null || val.isEmpty ? 'Bắt buộc' : null,
                        ),
                      ],
                    ),
                  ),
                ],
              ),

              const SizedBox(height: 48),

              // Nút Đăng
              SizedBox(
                height: 56,
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    backgroundColor: const Color(0xFF0F172A),
                    foregroundColor: Colors.white,
                    elevation: 0,
                    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
                  ),
                  onPressed: () {
                    if (_formKey.currentState!.validate() && _images.isNotEmpty) {
                      // TODO: Call API
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(content: Text('Đang xử lý tải lên...')),
                      );
                    } else if (_images.isEmpty) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(content: Text('Vui lòng chọn ít nhất 1 ảnh!')),
                      );
                    }
                  },
                  child: Text(
                    'ĐĂNG PHIÊN ĐẤU GIÁ',
                    style: GoogleFonts.inter(fontWeight: FontWeight.bold, letterSpacing: 1.2),
                  ),
                ),
              ),
              const SizedBox(height: 40),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildInputLabel(String text) {
    return Padding(
      padding: const EdgeInsets.only(bottom: 8.0),
      child: Text(
        text,
        style: GoogleFonts.inter(fontSize: 12, fontWeight: FontWeight.bold, color: const Color(0xFF0F172A)),
      ),
    );
  }

  InputDecoration _buildInputDecoration(String hint) {
    return InputDecoration(
      hintText: hint,
      hintStyle: GoogleFonts.inter(color: Colors.grey.shade400, fontSize: 14),
      filled: true,
      fillColor: Colors.grey.shade50,
      contentPadding: const EdgeInsets.all(16),
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
    );
  }
}
