package webScan

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func PortalScan(host string) {
	// outputFile := "portalScan.log"
	url := "https://" + host
	target, err := http.NewRequest(http.MethodGet, url, nil)
	defer func() {
		recover()
	}()
	if err != nil {
		panic(err)
	}

	target.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	// you can change your User-Agent here
	c := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	request, err := c.Do(target)
	if err != nil {
		fmt.Println(err)
	}
	status := request.Status
	defer request.Body.Close()
	// to check in automatically

	res, _ := io.ReadAll(request.Body)
	fmt.Println(status)
	fmt.Println(string(res))
}
