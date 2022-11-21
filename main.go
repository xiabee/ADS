package main

import (
	"ads/webScan"
)

func main() {
	// secretScan.KeyScan("input.txt")
	webScan.HttpsScan("input.txt")
	// fmt.Println("Process Finished!")
}
