package main

import (
	"fmt"
	"internal/config"
)

func main() {
	str := config.CreateConfig()
	fmt.Println(str)

	str = config.SetBlacklist("howdy")
	fmt.Println(str)

	str = config.SetSource("hello")
	fmt.Println(str)

	fmt.Println(config.GetBlacklist(), config.GetSource())
}
