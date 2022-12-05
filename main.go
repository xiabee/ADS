package main

import (
	"ads/webScan"
)

func main() {
	// secretScan.KeyScan("input.txt")
	// webScan.HttpsScan("src_domains-cloud.txt")
	webScan.SSHScan("src_ips.txt")
	//fmt.Println("Process Finished!")
}
