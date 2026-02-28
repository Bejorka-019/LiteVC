package main

import (
	"fmt"
	cmd "litevc/CMD"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please use command litevc --help for more information")
		return
	}

	command := strings.ToLower(os.Args[1])

	switch command {
	case "init":
		{
			fmt.Print("Init")
			cmd.Init()
		}
	default:
		{
			fmt.Println("Please use command litevc --help for more information")
		}
	}
}
