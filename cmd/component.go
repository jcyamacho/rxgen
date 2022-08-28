package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/caarlos0/log"
	"github.com/iancoleman/strcase"
	"github.com/jcyamacho/rxgen/generators/component"
	"github.com/spf13/cobra"
)

func init() {
	componentCmd.Flags().String("style", string(component.StyleCssMod), fmt.Sprintf("%v", component.Styles))
	componentCmd.Flags().String("styled-postfix", "styles", "Styled components file name postfix, pj: Component.postfix.ts")

	rootCmd.AddCommand(componentCmd)
}

var componentCmd = &cobra.Command{
	Use: "component <name>",
	Aliases: []string{
		"cmp",
	},
	Short: "Create React component",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		logger := log.New(os.Stderr)
		logger.Infof("creating component: %s", name)

		path, err := os.Getwd()
		if err != nil {
			logger.WithError(err).Fatal("error gtting working dir")
		}

		style, err := cmd.Flags().GetString("style")
		if err != nil {
			logger.WithError(err).Fatal("error getting style flag")
		}

		styledPostfix, err := cmd.Flags().GetString("styled-postfix")
		if err != nil {
			logger.WithError(err).Fatal("error getting styled-postfix flag")
		}

		ctx := context.Background()
		logger.IncreasePadding()
		ctx = log.NewContext(ctx, logger)

		if err := component.Create(ctx, path, &component.Options{
			Name:          strcase.ToCamel(name),
			Style:         component.Style(style),
			StyledPostfix: styledPostfix,
		}); err != nil {
			logger.DecreasePadding()
			logger.WithError(err).Fatal("error creating component")
		}
	},
	Args: cobra.ExactArgs(1),
}
