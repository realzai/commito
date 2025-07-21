package cmd

import (
	"fmt"
	"github.com/realzai/commito/internal/ai"
	"github.com/realzai/commito/internal/config"
	"github.com/realzai/commito/internal/utils"
	"github.com/spf13/cobra"
)

var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Answer a question based on the changes made",
	Run: func(cmd *cobra.Command, args []string) {
		question := args[0]

		if question == "" {
			fmt.Println("❌ Please provide a question to answer.")
			return
		}

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

		println("Staged changes detected. Asking AI to answer the question based on these changes...")

		msg, err := client.Ask(question, diff)
		if err != nil {
			fmt.Println("❌ AI error:", err)
			return
		}

		println("Answering a question based on the changes made...")
		println(msg)
	},
}

func init() {
	rootCmd.AddCommand(answerCmd)
}
