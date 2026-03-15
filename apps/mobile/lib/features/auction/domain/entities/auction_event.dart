import 'bid.dart';

class AuctionStateEvent {
  final int currentPrice;
  final int bidIncrement;
  final String status;
  final int endTime;
  final int viewerCount;
  final List<BidHistory> bidHistory;

  AuctionStateEvent({
    required this.currentPrice,
    required this.bidIncrement,
    required this.status,
    required this.endTime,
    required this.viewerCount,
    required this.bidHistory,
  });
}

class ViewerEvent {
  final int viewerCount;

  ViewerEvent({
    required this.viewerCount,
  });
}

class BidEvent {
  final String productId;
  final int price;
  final List<BidHistory> bidHistory;

  BidEvent({
    required this.productId,
    required this.price,
    required this.bidHistory,
  });
}