SELECT
  sku,
  name,
  stock
FROM
  products
WHERE 
  stock < $1
