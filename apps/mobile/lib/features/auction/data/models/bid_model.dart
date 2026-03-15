import '../../domain/entities/bid.dart';

class BidHistoryModel extends BidHistory {
  BidHistoryModel({
    required super.price,
    required super.bidderName,
    required super.createdAt,
  });

  factory BidHistoryModel.fromJson(Map<String, dynamic> json) {
    return BidHistoryModel(
      price: json['price'] is String
          ? int.tryParse(json['price']) ?? 0
          : (json['price'] as num?)?.toInt() ?? 0,
      bidderName: json['bidder_name'] as String? ?? 'Anonymous',
      createdAt: json['created_at'] != null
          ? DateTime.tryParse(json['created_at'] as String) ?? DateTime.now()
          : DateTime.now(),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'price': price,
      'bidder_name': bidderName,
      'created_at': createdAt.toIso8601String(),
    };
  }
}

class BidModel extends Bid {
  BidModel({
    required super.productId,
    required super.amount,
  });

  factory BidModel.fromJson(Map<String, dynamic> json) {
    return BidModel(
      productId: json['product_id'] as String? ?? '',
      amount: json['amount'] is String
          ? int.tryParse(json['amount']) ?? 0
          : (json['amount'] as num?)?.toInt() ?? 0,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'product_id': productId,
      'amount': amount,
    };
  }
}
