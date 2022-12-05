package webScan

import (
	"ads/lib"
	"net"
	"strconv"
	"sync"
	"time"
)

func SSHScan(filename string) {
	var wg sync.WaitGroup
	urlist := lib.ReadLines(filename)
	wg.Add(len(urlist))
	for i := range urlist {
		url := urlist[i]
		go func() {
			PortScan(url, 22)
			//scan SSH Port
			wg.Done()
		}()
	}
	wg.Wait()
}

func PortScan(host string, port int) {
	outputFile := "sshScan.log"
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), time.Second*10)
	// timeout: 10 second
	defer func() {
		recover()
	}()
	if err != nil {
		lib.Log(outputFile, "[-] ", host, ": ", err)
	} else {
		lib.Log(outputFile, "[+] Warning: ", host, ":22 is open")
	}
	conn.Close()
}
