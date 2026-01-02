/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"internal/config"
	"internal/util"

	"github.com/spf13/cobra"
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "shows the raw configuration settings",
	Long: `Configuration settings consist of key:value pairs such as source:<filename>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			util.Blert(cmd.Long, args)
			return
		}
		config.DisplayConfig()
	},
}

func init() {
	configCmd.AddCommand(displayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
