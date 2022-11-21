package webScan

import (
	"ads/lib"
	"crypto/tls"
	"strings"
	"sync"
	"time"
)

func HttpsScan(filename string) {
	var wg sync.WaitGroup
	urlist := lib.ReadLines(filename)
	wg.Add(len(urlist))
	for i := range urlist {
		url := urlist[i]
		go func() {
			TlsScan(url)
			wg.Done()
		}()
	}
	wg.Wait()
}

// use goroutine

func TlsScan(host string) {
	conn, err := tls.Dial("tcp", host+":443", nil)
	defer func() {
		recover()
	}()
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			// to check if the host is alive
			_, errhttp := tls.Dial("tcp", host+":80", nil)
			if errhttp != nil {
				lib.Log("httpsScan.log", "[+] Error: ", host, " ", errhttp)
			}
		} else {
			lib.Log("httpsScan.log", "[+] Error: ", host, " ", err)
		}
	}
	cert := conn.ConnectionState().PeerCertificates[0]
	currentTime := time.Now()
	crtUnix := currentTime.Unix()   // current time
	expUnix := cert.NotAfter.Unix() // expire time

	if expUnix-crtUnix < 0 {
		lib.Log("httpsScan.log", "[+] Expired: ", host)
	} else {
		lib.Log("httpsScan.log", "[-] Valid: ", host, cert.NotAfter)
	}
}
