class Product {
  final String id;
  final String sellerId;
  final String name;
  final String description;
  final int startingPrice;
  final int currentPrice;
  final int bidIncrement;
  final int status; // 0: Draft, 1: Active, 2: Completed, 3: Canceled
  final List<String> imageUrls;
  final DateTime startAt;
  final DateTime endAt;
  final String? winnerId;
  final DateTime createdAt;
  final DateTime updatedAt;

  const Product({
    required this.id,
    required this.sellerId,
    required this.name,
    required this.description,
    required this.startingPrice,
    required this.currentPrice,
    required this.bidIncrement,
    required this.status,
    required this.imageUrls,
    required this.startAt,
    required this.endAt,
    this.winnerId,
    required this.createdAt,
    required this.updatedAt,
  });

  /// Ảnh đại diện (ảnh đầu tiên), fallback nếu không có
  String get thumbnailUrl =>
      imageUrls.isNotEmpty ? imageUrls.first : '';

  /// Kiểm tra phiên đấu giá còn hoạt động
  bool get isActive => status == 1 && DateTime.now().isBefore(endAt);

  /// Thời gian còn lại
  Duration get timeRemaining {
    final remaining = endAt.difference(DateTime.now());
    return remaining.isNegative ? Duration.zero : remaining;
  }

  factory Product.fromJson(Map<String, dynamic> json) {
    // Helper to safely parse int from JSON (handles null, int, double)
    int safeInt(dynamic value) {
      if (value == null) return 0;
      if (value is int) return value;
      if (value is num) return value.toInt();
      return 0;
    }

    // Parse Go's uuid.NullUUID: {"UUID": "xxx", "Valid": true/false}
    String? parseWinnerId(dynamic value) {
      if (value == null) return null;
      if (value is String) return value.isEmpty ? null : value;
      if (value is Map) {
        final valid = value['Valid'] == true;
        if (valid && value['UUID'] != null) {
          final uuid = value['UUID'] as String;
          // UUID all zeros = no winner
          if (uuid == '00000000-0000-0000-0000-000000000000') return null;
          return uuid;
        }
      }
      return null;
    }

    return Product(
      id: json['id']?.toString() ?? '',
      sellerId: json['seller_id']?.toString() ?? '',
      name: json['name']?.toString() ?? '',
      description: json['description']?.toString() ?? '',
      startingPrice: safeInt(json['starting_price']),
      currentPrice: safeInt(json['current_price']),
      bidIncrement: safeInt(json['bid_increment']),
      status: safeInt(json['status']),
      imageUrls: json['image_urls'] != null
          ? List<String>.from(json['image_urls'])
          : [],
      startAt: json['start_at'] != null
          ? DateTime.parse(json['start_at'])
          : DateTime.now(),
      endAt: json['end_at'] != null
          ? DateTime.parse(json['end_at'])
          : DateTime.now(),
      winnerId: parseWinnerId(json['winner_id']),
      createdAt: json['created_at'] != null
          ? DateTime.parse(json['created_at'])
          : DateTime.now(),
      updatedAt: json['updated_at'] != null
          ? DateTime.parse(json['updated_at'])
          : DateTime.now(),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'seller_id': sellerId,
      'name': name,
      'description': description,
      'starting_price': startingPrice,
      'current_price': currentPrice,
      'bid_increment': bidIncrement,
      'status': status,
      'image_urls': imageUrls,
      'start_at': startAt.toIso8601String(),
      'end_at': endAt.toIso8601String(),
      'winner_id': winnerId,
      'created_at': createdAt.toIso8601String(),
      'updated_at': updatedAt.toIso8601String(),
    };
  }
}
