CREATE TABLE products (
  sku TEXT PRIMARY KEY,
  name TEXT,
  stock INT CHECK (stock >= 0),
  modified_at TIMESTAMP
)

CREATE TABLE stock_movements (
  id TEXT PRIMARY KEY,
  sku TEXT NOT NULL REFERENCES products(sku),
  movement INT NOT NUll,
  reason TEXT,
  created_at TIMESTAMP DEFAULT NOW()
)
