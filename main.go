package main

import (
	"ads/webScan"
	"fmt"
)

func main() {
	// secretScan.KeyScan("input.txt")
	webScan.HttpsScan("src_domains-cloud.txt")
	//webScan.PortalScan("tidbcloud.com")
	fmt.Println("Process Finished!")
}
