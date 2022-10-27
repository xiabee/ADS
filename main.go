package main

import (
	"ads/lib"
	"ads/secretScan"
	"time"
)

func main() {
	urlist := lib.ReadLines("input.txt")
	for _, url := range urlist {
		secretScan.Trufflehog(url)
	}
	time.Sleep(3 * time.Second)
}
