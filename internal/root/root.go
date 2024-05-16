package root

import (
	"github.com/llm-operator/cli/internal/auth"
	"github.com/llm-operator/cli/internal/chat"
	"github.com/llm-operator/cli/internal/configs"
	"github.com/llm-operator/cli/internal/files"
	"github.com/llm-operator/cli/internal/finetuning"
	"github.com/llm-operator/cli/internal/models"
	"github.com/llm-operator/cli/internal/ui"
	"github.com/llm-operator/cli/internal/version"
	"github.com/spf13/cobra"
)

// cmd represents the base command when called without any subcommands.
func cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "llmo",
		Short:              "LLM Operator CLI",
		DisableFlagParsing: false,
	}
	cmd.PersistentFlags().StringVar(&ui.Color, "color", string(ui.ColorAuto), "Control color output. Available options are 'auto', 'always' and 'never'.")

	cmd.AddCommand(auth.Cmd())
	cmd.AddCommand(chat.Cmd())
	cmd.AddCommand(configs.Cmd())
	cmd.AddCommand(files.Cmd())
	cmd.AddCommand(finetuning.Cmd())
	cmd.AddCommand(models.Cmd())
	cmd.AddCommand(version.Cmd())
	cmd.SilenceUsage = true

	return cmd
}

// Execute adds all child commands to the root command.
func Execute() error {
	return cmd().Execute()
}
