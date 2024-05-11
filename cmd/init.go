/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kidminks/gofold/internal"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project name]",
	Short: "Generate the base structure of your go project with the given name",
	Long: `init will create a new folder with project name with 
the appropriate folders and imports 
It requires a project name for generation for eg :-

gofold init fastDev`,
	Args: cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, "Please specify the path for the project")
		} else if len(args) == 1 {
			comps = cobra.AppendActiveHelp(comps, "Please specify the module")
		} else if len(args) == 2 {
			comps = cobra.AppendActiveHelp(comps, "This command does not take any more arguments (but may accept flags)")
		} else {
			comps = cobra.AppendActiveHelp(comps, "ERROR: Too many arguments specified")
		}
		return comps, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		module := args[1]
		configFile, _ := cmd.Flags().GetString("config")
		if configFile != "" {
			internal.GenerateStructureUsingConfigFile(path, configFile, false)
		} else {
			internal.GenerateDefaultConfigFile(path)
			internal.GenerateStructureUsingConfigFile(path, internal.DefaultConfigFile, true)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
