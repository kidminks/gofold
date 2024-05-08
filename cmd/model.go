/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/kidminks/gofold/internal"
	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model [name] [field:type]",
	Short: "Generate a model file with basic crud operation on model in the given file",
	Long: `model will create a new file in folder linked to model folder 
	It requires a model name and field:type pair for generation for eg :-
	
	gofold model User id:int64 name:string email:string password:string`,
	Args: cobra.MinimumNArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, "Please specify the model name")
		} else if len(args) == 1 {
			comps = cobra.AppendActiveHelp(comps, "Please specify the field:type for creating model")
		}
		return comps, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		fmt.Println(configFile)
		if configFile == "" {
			slog.Error("config file not specified")
			os.Exit(1)
		}
		internal.GenerateModel(args[0], configFile, args[1:])
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
