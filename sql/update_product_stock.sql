UPDATE products
SET stock = stock+$2
WHERE sku = $1 AND stock + $2 >= 0
RETURNING stock
