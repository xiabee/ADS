package webScan

import (
	"ads/lib"
	"crypto/tls"
	"strings"
	"sync"
	"time"
)

func HttpsScan(filename string) {
	var HttpsWg sync.WaitGroup
	urlist := lib.ReadLines(filename)
	HttpsWg.Add(len(urlist))
	for i := range urlist {
		url := urlist[i]
		go func() {
			TlsScan(url)
			HttpsWg.Done()
		}()
	}
	HttpsWg.Wait()
}

// use goroutine

func TlsScan(host string) {
	outputFile := "httpsScan.log"
	conn, err := tls.Dial("tcp", host+":443", new(tls.Config))

	defer func() {
		recover()
	}()
	// defer & recover to avoid crash
	if err != nil {
		if strings.Contains(err.Error(), "EOF") {
			lib.Log(outputFile, "Oops! EOF Error! ", host, err)
			TlsScan(host)
			// to avoid TLS EOF error
		}
		if strings.Contains(err.Error(), "connection refused") || strings.Contains(err.Error(), "operation timed out") {
			// to check if the host is alive
			_, errhttp := tls.Dial("tcp", host+":80", new(tls.Config))
			if errhttp != nil {
				lib.Log(outputFile, "[+] Error: ", host, " ", errhttp)
			} else {
				lib.Log(outputFile, "[+] TLS Warning: ", host, " is running")
			}
		} else {
			lib.Log(outputFile, "[+] Error: ", host, " ", err)
		}
	}

	cert := conn.ConnectionState().PeerCertificates[0]
	currentTime := time.Now()
	crtUnix := currentTime.Unix()   // current time
	expUnix := cert.NotAfter.Unix() // expire time

	if expUnix-crtUnix < 0 {
		lib.Log(outputFile, "[+] Expired: ", host)
	} else {
		lib.Log(outputFile, "[-] Valid: ", host, " ", cert.NotAfter)
	}
}
