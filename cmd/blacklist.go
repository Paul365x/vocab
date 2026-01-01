/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// blacklistCmd represents the blacklist command
var blacklistCmd = &cobra.Command{
	Use:   "blacklist",
	Short: "Add things to ignore when running a concordance",
	Long: `I originally wrote this to help translate the Vulgate. This had numbers, the word 
"Anonymous" and other things I wasn't interested in when running a concordance. This file loads
all those definitions, Each definition is by itself on a line in a text file:
- <number> will ignore all numbers;
- word

Call: vocab config blacklist <file/path>`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			blertBl(cmd, args)
			return
		}
		fmt.Println("Adding the file: ", args[0], " as the blacklist")
		// read in the config file and put this file in the blacklist position
	},
}

func blertBl(cmd *cobra.Command, args []string) {
	fmt.Println(cmd.Long)
	fmt.Println("You gave me: ", args)
}

func init() {
	configCmd.AddCommand(blacklistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blacklistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blacklistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
