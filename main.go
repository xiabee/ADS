package main

import (
	"ads/lib"
	"ads/secretScan"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	urlist := lib.ReadLines("input.txt")
	wg.Add(len(urlist))

	for _, url := range urlist {
		go func() {
			secretScan.Trufflehog(url)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Process Finished!")
}
