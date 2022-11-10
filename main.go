package main

import (
	"ads/webScan"
)

func main() {
	// secretScan.KeyScan("input.txt")
	webScan.HttpsScan("src_domains.txt")
	// fmt.Println("Process Finished!")
}
