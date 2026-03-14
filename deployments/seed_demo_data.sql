DO $$
DECLARE
    seller UUID := '7f16db01-676c-4fbb-94fa-4819ea3f8f68';
    
    p1  UUID; p2  UUID; p3  UUID; p4  UUID; p5  UUID;
    p6  UUID; p7  UUID; p8  UUID; p9  UUID; p10 UUID;
    p11 UUID; p12 UUID;
BEGIN

-- INSERT 12 PRODUCTS 
INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller, 
     'Rolex Submariner 1960 Vintage', 
     'Đồng hồ Rolex Submariner sản xuất năm 1960, tình trạng nguyên bản 95%. Có giấy chứng nhận từ Rolex Geneva. Một trong những chiếc Submariner hiếm nhất thế giới.',
     50000000, 127000000, 5000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408534/Gemini_Generated_Image_l8vw3kl8vw3kl8vw.png'],
     NOW() - INTERVAL '2 days', NOW() + INTERVAL '30 days')
RETURNING id INTO p1;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Tranh Sơn Dầu "Hoàng Hôn Sài Gòn" - Họa Sĩ Nguyễn Thanh',
     'Tác phẩm sơn dầu trên canvas 120x80cm, vẽ năm 2020. Phong cảnh hoàng hôn trên sông Sài Gòn với ánh nắng vàng rực rỡ. Có chữ ký và chứng nhận của họa sĩ.',
     15000000, 42000000, 2000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408519/Gemini_Generated_Image_d3uqp1d3uqp1d3uq.png'],
     NOW() - INTERVAL '1 day', NOW() + INTERVAL '30 days')
RETURNING id INTO p2;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Leica M3 Double Stroke 1954',
     'Máy ảnh Leica M3 phiên bản Double Stroke đời đầu sản xuất năm 1954. Ống kính Summicron 50mm f/2. Tình trạng hoạt động hoàn hảo, ngoại hình 90%.',
     30000000, 68000000, 3000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408538/Gemini_Generated_Image_h2e9sth2e9sth2e9.png'],
     NOW() - INTERVAL '3 days', NOW() + INTERVAL '30 days')
RETURNING id INTO p3;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Vespa Sprint 150 1967 Restored',
     'Vespa Sprint 150cc đời 1967 đã được phục chế toàn bộ. Sơn nguyên bản màu xanh Azzurro. Động cơ chạy êm, có giấy tờ đầy đủ.',
     25000000, 51000000, 2000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408710/Gemini_Generated_Image_pefj9zpefj9zpefj.png'],
     NOW() - INTERVAL '1 day', NOW() + INTERVAL '30 days')
RETURNING id INTO p4;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Gibson Les Paul Standard 1959 Sunburst',
     'Holy Grail của thế giới guitar! Gibson Les Paul Standard 1959 với finish Sunburst nguyên bản. Pickup PAF gốc. Đã được thẩm định bởi Gruhn Guitars Nashville.',
     200000000, 350000000, 10000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408747/Gemini_Generated_Image_r02jdhr02jdhr02j.png'],
     NOW() - INTERVAL '5 days', NOW() + INTERVAL '30 days')
RETURNING id INTO p5;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Sapphire Kashmir 8.5 Carat No Heat',
     'Viên Sapphire Kashmir 8.5 carat không qua xử lý nhiệt. Màu xanh "Cornflower Blue" đặc trưng. Chứng nhận GRS và Gübelin. Cực hiếm trên thị trường.',
     500000000, 720000000, 20000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408534/ChatGPT_Image_20_05_01_13_thg_3_2026.png'],
     NOW() - INTERVAL '2 days', NOW() + INTERVAL '30 days')
RETURNING id INTO p6;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Truyện Kiều Bản In 1866 - Nguyễn Du',
     'Bản in Truyện Kiều năm 1866, một trong những bản in sớm nhất còn sót lại. Giấy dó nguyên bản, chữ Nôm rõ nét. Đã qua kiểm định Viện Hán Nôm Việt Nam.',
     80000000, 95000000, 5000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408543/ChatGPT_Image_20_04_59_13_thg_3_2026.png'],
     NOW() - INTERVAL '1 day', NOW() + INTERVAL '30 days')
RETURNING id INTO p7;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Apple Watch Ultra 2 Hermès Edition',
     'Phiên bản giới hạn Apple Watch Ultra 2 hợp tác cùng Hermès. Dây da Swift Orange, mặt số Hermès độc quyền. Nguyên seal, fullbox.',
     35000000, 48000000, 1000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773409034/Gemini_Generated_Image_o4tq87o4tq87o4tq.png'],
     NOW(), NOW() + INTERVAL '30 days')
RETURNING id INTO p8;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Bình Gốm Hoa Lam Triều Lê Thế Kỷ 15',
     'Bình gốm hoa lam men trắng xanh thời Lê Sơ, thế kỷ 15. Hoa văn rồng mây tinh xảo. Đã trưng bày tại Bảo tàng Lịch sử Quốc gia. Có giấy xuất xứ đầy đủ.',
     150000000, 210000000, 10000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408536/ChatGPT_Image_20_04_57_13_thg_3_2026.png'],
     NOW() - INTERVAL '4 days', NOW() + INTERVAL '30 days')
RETURNING id INTO p9;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Nintendo Game Boy 1989 Sealed NIB',
     'Nintendo Game Boy đời đầu 1989 còn nguyên seal chưa bóc. Hộp nguyên vẹn 100%, VGA Graded 85+. Một trong những chiếc Game Boy sealed hiếm nhất thế giới.',
     20000000, 38000000, 2000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408535/Gemini_Generated_Image_hasnathasnathasn.png'],
     NOW() - INTERVAL '2 days', NOW() + INTERVAL '30 days')
RETURNING id INTO p10;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Hot Wheels Pink Rear-Loading Beach Bomb 1969',
     'Hot Wheels Beach Bomb phiên bản prototype màu hồng 1969. Đây là chiếc xe mô hình đắt nhất thế giới. Chỉ có 2 chiếc tồn tại. Đã xác thực bởi Bruce Pascal.',
     300000000, 450000000, 15000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408537/ChatGPT_Image_20_04_53_13_thg_3_2026.png'],
     NOW() - INTERVAL '3 days', NOW() + INTERVAL '30 days')
RETURNING id INTO p11;

INSERT INTO products (id, seller_id, name, description, starting_price, current_price, bid_increment, status, image_urls, start_at, end_at)
VALUES
    (gen_random_uuid(), seller,
     'Nike Air Mag 2016 "Back to the Future" Size 10',
     'Nike Air Mag 2016 tự động buộc dây, lấy cảm hứng từ phim Back to the Future II. Size US 10. Chỉ sản xuất 89 đôi trên toàn thế giới. Fullbox, chưa qua sử dụng.',
     400000000, 580000000, 20000000, 1,
     ARRAY['https://res.cloudinary.com/dsn8dnaud/image/upload/v1773408548/Gemini_Generated_Image_35onx335onx335on.png'],
     NOW() - INTERVAL '1 day', NOW() + INTERVAL '30 days')
RETURNING id INTO p12;

-- INSERT PRODUCT STATS
INSERT INTO product_stats (product_id, bid_count, view_count, favorite_count) VALUES
    (p11,  45, 12500, 320),   
    (p12, 38, 9800,  280),     
    (p5,  32, 8200,  250),   
    (p6,  28, 7500,  200),   
    (p1, 25, 6800,  180),   
    (p3,  22, 5500,  150),   
    (p9,  20, 4800,  130),   
    (p2,  18, 4200,  170),   
    (p4,  15, 3800,  120),   
    (p10, 12, 3200,  90),    
    (p8,  8,  2100,  60),    
    (p7,  5,  1500,  40);    

END $$;
