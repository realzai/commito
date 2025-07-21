package cmd

import "github.com/spf13/cobra"

var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Answer a question based on the changes made",
	Run: func(cmd *cobra.Command, args []string) {
		println("Answering a question based on the changes made...")
	},
}

func init() {
	rootCmd.AddCommand(answerCmd)
}
