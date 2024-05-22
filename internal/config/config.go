package config

import (
	"context"
	"fmt"

	"github.com/llm-operator/cli/internal/auth/org"
	"github.com/llm-operator/cli/internal/auth/project"
	"github.com/llm-operator/cli/internal/configs"
	"github.com/llm-operator/cli/internal/runtime"
	"github.com/spf13/cobra"
)

// Cmd returns a new config command.
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "config",
		Short:              "Config commands",
		Args:               cobra.NoArgs,
		DisableFlagParsing: true,
	}
	cmd.AddCommand(createCmd())
	cmd.AddCommand(setCmd())
	return cmd
}

func createCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "create",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return configs.CreateNewConfig()
		},
	}
}
func setCmd() *cobra.Command {
	return &cobra.Command{
		Use: "set",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("expected <key> <value>")
			}
			return set(cmd.Context(), args[0], args[1])
		},
	}
}

func set(ctx context.Context, key, value string) error {
	env, err := runtime.NewEnv(ctx)
	if err != nil {
		return err
	}

	switch key {
	case "organization":
		o, found, err := org.FindOrgByTitle(env, value)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("organization not found")
		}
		env.Config.Context.OrganizationID = o.Id
		if err := env.Config.Save(); err != nil {
			return err
		}
	case "project":
		p, found, err := project.FindProjectByTitle(env, value, "")
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("project not found")
		}
		env.Config.Context.ProjectID = p.Id
		if err := env.Config.Save(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown key %s", key)
	}

	return nil
}
