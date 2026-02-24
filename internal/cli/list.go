package cli

import (
	"database/sql"
	"log"

	"github.com/spf13/cobra"

	"mcbulazs/mf-take-home-task/internal/db"
	"mcbulazs/mf-take-home-task/internal/product"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all products",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.ConnectToDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		repo := product.NewSQLProductRepository(db)
		service := product.NewProductService(repo)
		err = service.List()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func listProducts(db *sql.DB) error {
	return nil
}
