CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price NUMERIC(10,2),
    stock INT
);

INSERT INTO products (name, price, stock)
VALUES
('Apple', 1.20, 100),
('Banana', 0.80, 50),
('Orange', 0.80, 200);
