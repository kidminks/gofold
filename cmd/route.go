/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/kidminks/gofold/internal"
	"github.com/spf13/cobra"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route [model_name]",
	Short: "Generate a route file with basic CRUD operations",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

handler will create a new file in folder linked to handler folder 
	It requires a model name for generation for eg :-
	
	gofold route User`,
	Args: cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, "Please specify the model name")
		} else if len(args) == 1 {
			comps = cobra.AppendActiveHelp(comps, "This command does not take any more arguments (but may accept flags)")
		} else {
			comps = cobra.AppendActiveHelp(comps, "ERROR: Too many arguments specified")
		}
		return comps, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		if configFile == "" {
			slog.Error("config file not specified")
			os.Exit(1)
		}
		hErr := internal.GenerateHandler(args[0], configFile)
		if hErr == nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(routeCmd)
}
