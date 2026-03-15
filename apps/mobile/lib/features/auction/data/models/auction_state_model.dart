import '../../domain/entities/auction_event.dart';
import '../../domain/entities/bid.dart';
import 'bid_model.dart';

class AuctionStateEventModel extends AuctionStateEvent {
  AuctionStateEventModel({
    required super.currentPrice,
    required super.bidIncrement,
    required super.status,
    required super.endTime,
    required super.viewerCount,
    required super.bidHistory,
  });

  factory AuctionStateEventModel.fromJson(Map<String, dynamic> json) {
    int _parseInt(dynamic value) {
      if (value is int) return value;
      if (value is String) return int.tryParse(value) ?? 0;
      return 0;
    }

    return AuctionStateEventModel(
      currentPrice: _parseInt(json['current_price']),
      bidIncrement: _parseInt(json['bid_increment']),
      status: json['status'] as String? ?? 'active',
      endTime: _parseInt(json['end_time']),
      viewerCount: _parseInt(json['viewer_count']),
      bidHistory: (json['bid_history'] as List<dynamic>?)
              ?.map((e) => BidHistoryModel.fromJson(e as Map<String, dynamic>))
              .toList() ??
          [],
    );
  }
}

class ViewerEventModel extends ViewerEvent {
  ViewerEventModel({
    required super.viewerCount,
  });

  factory ViewerEventModel.fromJson(dynamic json) {
    int count = 0;
    if (json is num) {
      count = json.toInt();
    } else if (json is String) {
      count = int.tryParse(json) ?? 0;
    }
    return ViewerEventModel(viewerCount: count);
  }
}

class BidEventModel extends BidEvent {
  BidEventModel({
    required super.productId,
    required super.price,
    required super.bidHistory,
  });

  factory BidEventModel.fromJson(Map<String, dynamic> json, String productId) {
    int _parseInt(dynamic value) {
      if (value is int) return value;
      if (value is String) return int.tryParse(value) ?? 0;
      return 0;
    }

    return BidEventModel(
      productId: productId,
      price: _parseInt(json['price']),
      bidHistory: (json['bid_history'] as List<dynamic>?)
              ?.map((e) => BidHistoryModel.fromJson(e as Map<String, dynamic>))
              .toList() ??
          [],
    );
  }
}
