DROP TRIGGER IF EXISTS trg_copy_favorite_count ON product_favorites;
DROP FUNCTION IF EXISTS update_favorite_count();
DROP TABLE IF EXISTS product_favorites;