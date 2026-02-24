SELECT
  1
FROM
  stock_movements
WHERE
  id = $1
  AND sku = $2
  AND movement = $3
