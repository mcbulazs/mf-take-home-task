package sql

import _ "embed"

//go:embed migrate_schema.sql
var MigrateSchemaSQL string

//go:embed migrate_seed.sql
var MigrateSeedSQL string

//go:embed list_products.sql
var ListProductsSQL string
