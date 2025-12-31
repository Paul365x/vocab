package main

import (
	"os"
	"sort"
	"strconv"

	"fmt"
	"text/scanner"

	"github.com/dariubs/percent"
)

func main() {
	conc := make(map[string]int)
	var s scanner.Scanner
	f, err := os.Open("vul.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s.Init(f)
	var keys []string

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		word := s.TokenText()

		// weed the stream
		_, err := strconv.Atoi(word)
		if err == nil {
			continue
		}
		switch word {
		case "Anonymous", "Latin", "Bible", "(Vul)", ",", "(", ")", "", ".", ":", ";", "%":
			continue

		}
		_, ok := conc[word]
		if ok {
			conc[word]++
			/*
				if word == "et" {
					fmt.Printf("%s: %d, %d\n", word, conc[word], val)
				}*/
		} else {
			/* if word == "et" {
				fmt.Println("first et")
			}*/
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
	for key := range keys {
		count := conc[keys[key]]
		acc += count
		//order++
		fmt.Printf("%d %s %d | %.2f%% %.2f%%\n", key, keys[key], conc[keys[key]],
			percent.PercentOf(count, tot), percent.PercentOf(acc, tot))
	}

}
