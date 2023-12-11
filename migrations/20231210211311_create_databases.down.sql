DROP INDEX IF EXISTS idx_product_unique_code;
DROP INDEX IF EXISTS idx_reservation_product_id;
DROP INDEX IF EXISTS idx_product_store_product_id;
DROP INDEX IF EXISTS idx_product_store_store_id;

DROP TABLE IF EXISTS product_store;
DROP TABLE IF EXISTS reservation;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS store;