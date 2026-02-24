package sql

import _ "embed"

//go:embed migrate_schema.sql
var MigrateSchemaSQL string

//go:embed migrate_seed.sql
var MigrateSeedSQL string

//go:embed list_products.sql
var ListProductsSQL string

//go:embed sum_product_stock.sql
var SumProductStockSQL string

//go:embed count_products.sql
var CountProductsSQL string

//go:embed list_top_products.sql
var TopNProductsSQL string

//go:embed list_low_stock_products.sql
var LowestProductsUnderSQL string

//go:embed insert_stock_movement.sql
var InsertStockMovementSQL string

//go:embed update_product_stock.sql
var UpdateProductStockSQL string

//go:embed get_stock_movement.sql
var GetStockMovementSQL string
