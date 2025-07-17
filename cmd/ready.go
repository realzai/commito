package cmd

import (
	"fmt"

	"github.com/realzai/commito/internal/utils"
	"github.com/spf13/cobra"
)

var readyCmd = &cobra.Command{
	Use:   "ready",
	Short: "Check if commito is ready and show your configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := utils.EnsureConfigured()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("✅ Commito is ready to go!")
		fmt.Println("Provider:", cfg.Provider)
		fmt.Println("Model:   ", cfg.Model)
	},
}

func init() {
	rootCmd.AddCommand(readyCmd)
}
