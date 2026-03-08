CREATE TABLE product_favorites (
    user_id UUID NOT NULL,
    product_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    
    PRIMARY KEY (user_id, product_id),
    
    CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX idx_favorites_user_id ON product_favorites(user_id);

CREATE OR REPLACE FUNCTION update_favorite_count()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        UPDATE product_stats 
        SET favorite_count = favorite_count + 1, updated_at = NOW()
        WHERE product_id = NEW.product_id;
    ELSIF (TG_OP = 'DELETE') THEN
        UPDATE product_stats 
        SET favorite_count = favorite_count - 1, updated_at = NOW()
        WHERE product_id = OLD.product_id;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_copy_favorite_count
AFTER INSERT OR DELETE ON product_favorites
FOR EACH ROW EXECUTE FUNCTION update_favorite_count();