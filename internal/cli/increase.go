package cli

import (
	"log"

	"github.com/spf13/cobra"

	"mcbulazs/mf-take-home-task/internal/db"
	"mcbulazs/mf-take-home-task/internal/services"
)

var increaseCmd = &cobra.Command{
	Use:   "increase",
	Short: "Increase stock for a product",
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
		err = service.IncreseQuantity(id, sku, qty, reason)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	increaseCmd.Flags().String("id", "", "Idempotency key (required)")
	increaseCmd.Flags().String("sku", "", "Product's sku (required)")
	increaseCmd.Flags().Int("qty", 0, "The quantity (required)")
	increaseCmd.Flags().String("reason", "", "The reason for the increase")
}
