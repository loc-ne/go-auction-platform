CREATE TABLE product_stats (
    product_id UUID PRIMARY KEY, 
    bid_count INTEGER DEFAULT 0,
    view_count INTEGER DEFAULT 0,
    favorite_count INTEGER DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW(),
    
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
);