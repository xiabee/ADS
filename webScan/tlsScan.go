package webScan

import (
	"ads/lib"
	"crypto/tls"
	"fmt"
	"time"
)

func HttpsScan(filename string) {
	urlist := lib.ReadLines(filename)
	for _, url := range urlist {
		TlsScan(url)
	}
}

func TlsScan(host string) {
	conn, err := tls.Dial("tcp", host+":443", nil)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[!] Error: ", err)
		}
	}()
	if err != nil {
		fmt.Println("[+] Error: ", host, " ", err)
		lib.Log("httpsScan.log", "[+] Error: ", host, " ", err)
	}
	cert := conn.ConnectionState().PeerCertificates[0]
	currentTime := time.Now()
	crtUnix := currentTime.Unix()   // current time
	expUnix := cert.NotAfter.Unix() // expire time
	if expUnix-crtUnix < 0 {
		fmt.Println("[+] Expired: ", host)
		lib.Log("httpsScan.log", "[+] Expired: ", host)
	} else {
		fmt.Println("[-] Valid: ", host, cert.NotAfter)
		lib.Log("httpsScan.log", "[-] Valid: ", host, cert.NotAfter)
	}
}
