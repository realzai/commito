package cmd

import "github.com/spf13/cobra"

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest a commit message based on the changes made",
	Run: func(cmd *cobra.Command, args []string) {
		println("Suggesting a commit message based on the changes made...")
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}
