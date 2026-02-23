package cli

import (
	"log"

	"github.com/spf13/cobra"

	"mcbulazs/mf-take-home-task/internal/db"
	"mcbulazs/mf-take-home-task/internal/services"
)

var decreaseCmd = &cobra.Command{
	Use:   "decrease",
	Short: "Decrease stock for a product",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		sku, _ := cmd.Flags().GetString("sku")
		qty, _ := cmd.Flags().GetInt("qty")
		reason, _ := cmd.Flags().GetString("reason")
		if id == "" || sku == "" || qty <= 0 {
			log.Fatal("Required flags: --id, --sku, --qty>0")
		}

		db, err := db.ConnectToDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		service := services.NewProductService(db)
		err = service.DecreaseQuantity(id, sku, qty, reason)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	decreaseCmd.Flags().String("id", "", "Idempotency key (required)")
	decreaseCmd.Flags().String("sku", "", "Product's sku (required)")
	decreaseCmd.Flags().Int("qty", 0, "The quantity (required)")
	decreaseCmd.Flags().String("reason", "", "The reason for the increase")
}
