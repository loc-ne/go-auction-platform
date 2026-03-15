class WsMessageModel {
  final String roomId;
  final String userId;
  final String action;
  final dynamic payload;

  WsMessageModel({
    required this.roomId,
    required this.userId,
    required this.action,
    required this.payload,
  });

  factory WsMessageModel.fromJson(Map<String, dynamic> json) {
    return WsMessageModel(
      roomId: json['roomId'] as String? ?? '',
      userId: json['userId'] as String? ?? '',
      action: json['action'] as String? ?? '',
      payload: json['payload'],
    );
  }
}
