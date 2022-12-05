# ADS
Active Defense System for cyber attacks.

## Basic Informaition

* Author: [xiabee](https://github.com/xiabee),  [PingCAP](https://github.com/pingcap) Security Team
* Update time: 2022.12.5

* Version: 1.2



## Usage

In `main.go`, update the code with the correct input files:

```go
func main() {
	// secretScan.KeyScan("input.txt")
	webScan.HttpsScan("src_domains-cloud.txt")
	fmt.Println("Process Finished!")
}
```

Then go run or go build.



## Current Fuction

* Leaked key scan in GitHub
* TLS scan
* SSH scan



## Changelog

* Add goroutine
* Add TLS Scanning



## Ongoing

* Multi-func design

