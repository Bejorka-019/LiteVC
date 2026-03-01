package main

import (
	"fmt"
	cmd "litevc/cmd"
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
	case "detect":
		{
			fmt.Println("Scan file and add changes to staging area")
			cmd.Detect()
		}
	default:
		{
			fmt.Println("Please use command litevc --help for more information")
		}
	}
}
