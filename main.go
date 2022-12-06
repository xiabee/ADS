package main

import (
	"ads/lib"
	"ads/secretScan"
	"ads/webScan"
	"fmt"
	"os"
)

func main() {
	file := ""
	args := os.Args
	argc := len(args)
	if argc == 1 {
		lib.Usage()
		return
	}

	for i := range args {
		if args[i] == "-h" || args[i] == "--h" || args[i] == "-help" || args[i] == "--help" {
			lib.Help()
			return
		}
		if (args[i] == "-f" || args[i] == "--f" || args[i] == "-file" || args[i] == "--file") && i < argc-1 {
			file = args[i+1] // check input file
		}
	}
	if file == "" {
		fmt.Println("Use \"-f\" to specify input file")
		return
	}

	for i := range args {
		switch args[i] {
		case "-ssh", "--ssh":
			webScan.SSHScan(file)
			fmt.Println("The result is in sshScan.log")
		case "-https", "--https":
			webScan.HttpsScan(file)
			fmt.Println("The result is in httpsScan.log")
		case "-key", "--key":
			secretScan.KeyScan(file)
			fmt.Println("The result is in secretScan.log")
		}
	}
}
