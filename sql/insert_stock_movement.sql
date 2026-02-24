INSERT INTO stock_movements(id,
sku,
movement,
reason,
created_at)
VALUES($1,
$2,
$3,
$4,
NOW())
