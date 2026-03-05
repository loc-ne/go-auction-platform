CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    seller_id UUID NOT NULL,            
    name VARCHAR(255) NOT NULL,
    description TEXT,
    starting_price BIGINT NOT NULL,   
    current_price BIGINT NOT NULL,    
    bid_increment BIGINT DEFAULT 0,   
    status INT DEFAULT 0,             -- 0: Draft, 1: Active, 2: Completed, 3: Canceled
    image_urls TEXT[],                  
    start_at TIMESTAMP NOT NULL,        
    end_at TIMESTAMP NOT NULL,        
    winner_id UUID NULL,              
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_products_status ON products(status);
CREATE INDEX idx_products_seller_id ON products(seller_id);
CREATE INDEX idx_products_winner_id ON products(winner_id);
CREATE INDEX idx_products_active_end_at ON products(status, end_at);
CREATE INDEX idx_products_status_price ON products(status, current_price);