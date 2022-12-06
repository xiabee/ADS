# ADS
Active Defense System for cyber attacks.

## Basic Informaition

* Author: [xiabee](https://github.com/xiabee),  [PingCAP](https://github.com/pingcap) Security Team
* Update time: 2022.12.6

* Version: 1.3



## Usage

### Clone and build

```bash
git clone https://github.com/xiabee/ADS
cd ADS
go build
```



### Execute

```bash
./ads -file to to specify input file
./ads -ssh to detect ssh ports
./ads -https to check tls certificate
./ads -key to scan leaked keys in github repo
```



### Examples





## Current Fuction

* Leaked key scan in GitHub
* TLS scan
* SSH scan



## Changelog

* Add goroutine
* Add TLS Scanning



## Ongoing

* Multi-func design

