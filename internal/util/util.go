package util

import (
	"fmt"
)

func Blert(msg string, args []string) {
	fmt.Println(msg)
	if len(args) > 0 {
		fmt.Println("You gave me: ", args)
	}
	
}