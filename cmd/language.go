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

// languageCmd represents the language command
var languageCmd = &cobra.Command{
	Use:   "language",
	Short: "Sets the language transform used in tranlation",
	Long: `We are using https://glosbe.com to do our translation. Because of attacks they have withdrawn their interface,
	so we are forced to use web scraping instead. This is fragile and will break whenever they make changes. To use this web scraping
	you need to supply the language transform. If you go to glosbe.com and do a lookup of a word in the address bar you will see something
	like "glosbe.com/la/en/quod. In this example we are transforming latin "quod" to english. Note the /la/en/ that is the string we want.
	
	Call: vocab config language <transform string>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			util.Blert(cmd.Long, args)
			return
		}
		str := config.SetLanguage(args[0])
		if str["language"] == args[0] {
			fmt.Println("Added the language transform: ", args[0], " for all translation")
		}		
	},
}

func init() {
	configCmd.AddCommand(languageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// languageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// languageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
