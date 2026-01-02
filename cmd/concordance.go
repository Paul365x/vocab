/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"internal/config"
	"internal/util"
	"text/scanner"

	"github.com/dariubs/percent"

	"github.com/spf13/cobra"
)

// concordance globals
var blacklist map[string]int	// lookup table for blacklisted words
var source string
var chapters map[string]int	    // lookup table for chapter break words
var blNumber bool               // blacklist all numbers - set on the token <number> in the black list file
// concordanceCmd represents the concordance command
var concordanceCmd = &cobra.Command{
	Use:   "concordance",
	Short: "Takes a document and lists a frequency distribution of words",
	Long: `Will read a file and list the words used in order of frequency - most frequent to 
	lease frequent. It depends on configured:
	- blacklist file listing one word/symbol/character per line of things to ignore. Also supports <number>
	- source file is the document to read
	- chapters - if the -c flag is set, will output in chapter format based on the chapter file. Otherwise
	any chapter file will be added to the blacklist`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			util.Blert(cmd.Long, args)
			return
		}
		err := readConfig()
		if err != nil {
			panic(err)
		}
		scan()
	},
}

func init() {
	rootCmd.AddCommand(concordanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// concordanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// concordanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	
}

func readConfig() error {
	var err error
	// get our config
	source = config.GetSource();
	if source == "none" {
		return errors.New("You don't have a source file set. Set one with:\nvocab config source <file/path>\n")
	} 
	blFile := config.GetBlacklist()
	if blFile != "none" {
		blacklist, err = readList(blFile)
		if err != nil {
			return err
		}			
	} else {
		blacklist = make(map[string]int)
	}
	_,num := blacklist["<number>"]
	if num {
		blNumber = true
	} else {
		blNumber = false
	}
	chFile := config.GetChapters()	
	if chFile != "none" {
		chapters, err = readList(blFile)	
		if err != nil {
			return err
		}	
	} else {
		chapters = make(map[string]int)
	}
	return nil
}

func readList(file string) (map[string]int, error) {
	tmp := make(map[string]int)
	c, err := os.ReadFile(file)
	if err != nil {
		return tmp, err
	}
	lns := strings.Split(string(c), "\n")
	for ln := range lns {
		str := strings.Fields(lns[ln])
		if len(str) != 1 {
			return make(map[string]int), errors.New("Corrupt file: " + file + "  More than one word per line")
		}		
		lnNoW := strings.ReplaceAll(lns[ln], " ", "")
		if len(lnNoW) <= 0 {
			continue
		}
		tmp[lnNoW] = 1
	}
	return tmp, nil
} 

func scan() {
	conc := make(map[string]int)
	var s scanner.Scanner
	f, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s.Init(f)
	var keys []string

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		word := s.TokenText()
        if word == "\n" {
			continue
		}

		// weed the stream
		if blNumber {
		    _, err := strconv.Atoi(word)
		    if err == nil {
			    continue
		    }
	    } 
		_, blacklisted := blacklist[word]  
		if blacklisted {
			continue
		}
		
		_, ok := conc[word]
		if ok {
			conc[word]++
		} else {
			keys = append(keys, word)
			conc[word] = 1
		}
	}

	var tot int
	for _, v := range conc {
		tot += v
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return conc[keys[i]] > conc[keys[j]]
	})

	acc := 0
	//order := 0
	fmt.Printf("%-10s %-20s %-10s | %-5s %-5s\n","V_Size", "Word", "Count","%Tot","Acc %Tot")
	for key := range keys {
		count := conc[keys[key]]
		acc += count
		//order++
		fmt.Printf("%-10d %-20s %-10d | %5.2f%% %-5.2f%%\n", key, keys[key], conc[keys[key]],
			percent.PercentOf(count, tot), percent.PercentOf(acc, tot))
	}

}
