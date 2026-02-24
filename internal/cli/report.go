package cli

import (
	"log"

	"github.com/spf13/cobra"

	"mcbulazs/mf-take-home-task/internal/db"
	"mcbulazs/mf-take-home-task/internal/product"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Print a stock summary report for all products",
	Long:  "Generates a simple CLI report showing total SKU count, total stock units, top N products by stock, and low stock items based on threshold.",
	Run: func(cmd *cobra.Command, args []string) {
		top, _ := cmd.Flags().GetInt("top")
		lowStock, _ := cmd.Flags().GetInt("low-stock")

		db, err := db.ConnectToDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		repo := product.NewSQLProductRepository(db)
		service := product.NewProductService(repo)
		err = service.Report(top, lowStock)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	reportCmd.Flags().Int("top", 5, "Number of top products to display by stock (default 5)")
	reportCmd.Flags().Int("low-stock", 10, "Threshold for low stock products; all products with stock <= this value will be listed (default 10)")
}
