package cmd

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/jcyamacho/rxgen/generators/component"
	"github.com/jcyamacho/rxgen/internal/funcutil"
	"github.com/spf13/cobra"
)

func init() {
	componentCmd.Flags().String("style", string(component.StyleCssMod), fmt.Sprintf("%v", component.Styles))

	rootCmd.AddCommand(componentCmd)
}

var componentCmd = &cobra.Command{
	Use:   "component <name>",
	Short: "Create React component",
	Run: funcutil.Fatal2(func(cmd *cobra.Command, args []string) error {
		name := args[0]

		path, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error gtting working dir: %w", err)
		}

		style, err := cmd.Flags().GetString("style")
		if err != nil {
			return fmt.Errorf("error getting style flag: %w", err)
		}

		return component.Create(path, &component.Options{
			Name:  strcase.ToCamel(name),
			Style: component.Style(style),
		})
	}),
	Args: cobra.ExactArgs(1),
}
