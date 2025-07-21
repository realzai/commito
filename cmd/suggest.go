package cmd

import (
	"fmt"
	"github.com/realzai/commito/internal/ai"
	"github.com/realzai/commito/internal/config"
	"github.com/realzai/commito/internal/utils"
	"github.com/spf13/cobra"
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest a commit message based on the changes made",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := config.LoadConfig()
		client, err := ai.NewClientFromConfig(cfg)

		if err != nil {
			println("Error creating AI client:", err.Error())
			return
		}

		diff, _ := utils.GetStagedDiff()

		if diff == "" {
			fmt.Println("❌ No staged changes found. Please stage your changes before asking a question.")
			return
		}

		println("Staged changes detected. Asking AI to suggest a commit message based on these changes...")

		msg, err := client.Suggest(diff)
		if err != nil {
			println("❌ AI error:", err)
			return
		}
		println("Suggested commit message based on the changes made:")
		println(msg)
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
