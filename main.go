package main

import (
	"ads/webScan"
)

func main() {
	// secretScan.KeyScan("input.txt")
	webScan.HttpsScan("src_domains-cloud.txt")
	// fmt.Println("Process Finished!")
}
