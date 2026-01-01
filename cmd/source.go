/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Sets the source document",
	Long: `This sets the file or source document that is used by all other commands
	Call: vocab config source <file/path>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			blertSrc(cmd, args)
			return
		}
		fmt.Println("Adding the file: ", args[0], " as the source document")
		// read in the config file and put this file in the blacklist position
	},
}

func blertSrc(cmd *cobra.Command, args []string) {
	fmt.Println(cmd.Long)
	fmt.Println("You gave me: ", args)
}

func init() {
	configCmd.AddCommand(sourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
