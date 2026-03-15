class Bid {
    final String productId;
    final int amount;

    Bid({
        required this.productId,
        required this.amount,
    });
}

class BidHistory {
    final int price;
    final String bidderName;
    final DateTime createdAt;
    
    BidHistory({
        required this.price,
        required this.bidderName,
        required this.createdAt,
    });
}