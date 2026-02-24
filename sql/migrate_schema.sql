CREATE TABLE IF NOT EXISTS products(
  sku TEXT PRIMARY KEY,
  name TEXT,
  stock INT CHECK(stock >= 0),
  modified_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS stock_movements(
  id TEXT,
  sku TEXT NOT NULL REFERENCES products(sku),
  movement INT NOT NULL,
  reason TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  PRIMARY KEY(
    id,
    sku,
    movement
  )
)
