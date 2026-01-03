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

// chaptersCmd represents the chapters command
var chaptersCmd = &cobra.Command{
	Use:   "chapters",
	Short: "Add the markers for chapter output of concordance",
	Long: `Many documents are broken into parts or chapters. In the bible, 
	we have the books - genesis, exodus.. etc. By supplying key words that define the chapters, 
	concordance can output a separate list for each "chapter". The chapter list is a text file with 
	one key word per line
	
	Call: vocab config chapters <file/path>`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				util.Blert(cmd.Long, args)
				return
			}
			fmt.Println("Adding the file: ", args[0], " as the chapter list")
			config.SetChapters(args[0])
		},
	}
	
	
func init() {
	configCmd.AddCommand(chaptersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chaptersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chaptersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
