import '../entities/bid.dart';

abstract class AuctionRepository {
  Future<void> bid(String productId, int amount);
}