CREATE TABLE IF NOT EXISTS store (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    available BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size VARCHAR(50) NOT NULL,
    unique_code VARCHAR(50) UNIQUE NOT NULL,
    count INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS store_product (
    id SERIAL PRIMARY KEY,
    store_id INTEGER REFERENCES Store(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES Product(id) ON DELETE CASCADE,
    count INTEGER NOT NULL,
    UNIQUE (store_id, product_id)
);

CREATE TABLE IF NOT EXISTS reservation (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES Product(id) ON DELETE CASCADE,
    count INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_product_unique_code ON product(unique_code);
CREATE INDEX IF NOT EXISTS idx_store_product_store_id ON store_product(store_id);
CREATE INDEX IF NOT EXISTS idx_store_product_product_id ON store_product(product_id);
CREATE INDEX IF NOT EXISTS idx_reservation_product_id ON reservation(product_id);

INSERT INTO store (name, available) VALUES
    ('Store1', true),
    ('Store2', true),
    ('Store3', false),
    ('Store4', true),
    ('Store5', false),
    ('Store6', true),
    ('Store7', false),
    ('Store8', true),
    ('Store9', true),
    ('Store10', false);

INSERT INTO product (name, size, unique_code, count) VALUES
    ('Product1', 'Small', 'ABC123', 100),
    ('Product2', 'Medium', 'DEF456', 150),
    ('Product3', 'Large', 'GHI789', 200),
    ('Product4', 'Extra Large', 'JKL012', 120),
    ('Product5', 'Small', 'MNO345', 180),
    ('Product6', 'Large', 'PQR678', 90),
    ('Product7', 'Small', 'STU901', 120),
    ('Product8', 'Medium', 'VWX234', 80),
    ('Product9', 'Extra Large', 'YZA567', 110),
    ('Product10', 'Small', 'BCD890', 150);

INSERT INTO store_product (store_id, product_id, count) VALUES
    (1, 1, 50),
    (1, 2, 75),
    (2, 3, 30),
    (3, 4, 40),
    (4, 5, 60),
    (5, 6, 40),
    (4, 7, 60),
    (3, 8, 25),
    (2, 9, 30),
    (1, 10, 50);