package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/realzai/commito/internal/config"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure AI providers and other settings",
	Run: func(cmd *cobra.Command, args []string) {
		var provider, apiKey, model string

		err := survey.AskOne(&survey.Select{
			Message: "Choose an AI provider:",
			Options: []string{"OpenAI", "Grok"},
			Default: "OpenAI",
		}, &provider)
		if err != nil {
			fmt.Println("❌ Error selecting provider:", err)
			os.Exit(1)
		}

		err = survey.AskOne(&survey.Input{
			Message: fmt.Sprintf("Enter your %s API key:", provider),
		}, &apiKey, survey.WithValidator(survey.Required))
		if err != nil {
			fmt.Println("❌ Error entering API key:", err)
			os.Exit(1)
		}

		if provider == "Grok" {
			err = survey.AskOne(&survey.Input{
				Message: "Enter Grok model name:",
			}, &model, survey.WithValidator(survey.Required))
			if err != nil {
				fmt.Println("❌ Error entering model name:", err)
				os.Exit(1)
			}
		}

		cfg := config.Config{
			Provider: provider,
			ApiKey:   apiKey,
			Model:    model,
		}

		if err := config.SaveConfig(cfg); err != nil {
			fmt.Println("❌ Failed to save config:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Configuration saved successfully.")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
