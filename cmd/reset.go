/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"internal/config"
	"internal/util"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Empties the config file",
	Long: `Sets all config settings to default state. Note: you will need to at least set:
	\t- for concordance: source

	Call: vocab config reset
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			util.Blert(cmd.Long, args)
			return
		}		
		fmt.Println("Resetting the configuration file")
		config.CreateConfig()
	},
}

func init() {
	configCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
