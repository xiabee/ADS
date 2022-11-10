package secretScan

import (
	"ads/lib"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func KeyScan(filename string) {
	var wg sync.WaitGroup
	urlist := lib.ReadLines(filename)
	wg.Add(len(urlist))

	for _, url := range urlist {
		go func() {
			Trufflehog(url)
			// To scan leaked keys
			fmt.Println(url, " Done!")
			wg.Done()
		}()
	}
	wg.Wait()
}

func Trufflehog(url string) {
	// target repository
	cmd := exec.Command("trufflehog", "git", url, "--json")
	// To check by json mode

	type GIT struct {
		Commit     string `json:"commit"`
		File       string `json:"file"`
		Email      string `json:"email"`
		Repository string `json:"repository"`
		Timestamp  string `json:"timestamp"`
		Line       int    `json:"line"`
	}

	type DATA struct {
		Git GIT `json:"Git"`
	}

	type SOURCE struct {
		Data DATA `json:"Data"`
	}
	type RESP struct {
		SourceMetadata SOURCE `json:"SourceMetadata"`
		SourceID       int    `json:"SourceID"`
		SourceType     int    `json:"SourceType"`
		DetectorType   int    `json:"DetectorType"`
		DetectorName   string `json:"DetectorName"`
		DecoderName    string `json:"DecoderName"`
		Verified       bool   `json:"Verified"`
		Raw            string `json:"Raw"`
		Redacted       string `json:"Redacted"`
		StructuredData string `json:"StructuredData"`
	}
	// The json format of trufflehog's responses

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // standard output
	cmd.Stderr = &stderr // standard error output
	err := cmd.Run()
	if err != nil {
		log.Fatal("Trufflehog Execution error:  ", err)
	}
	if stderr.Len() != 0 {
		log.Fatal("Stderr: ", string(stderr.Bytes()))
	}
	// Error output
	// fmt.Println(string(stdout.Bytes()))

	res := bytes.Split(stdout.Bytes(), []byte("\n"))
	// Use \r\n to split the multi json, maybe unstable when running on the other platform
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	// cut off the last null element

	var resp []RESP
	var tmp RESP
	for _, value := range res {
		err = json.Unmarshal(value, &tmp)
		if err != nil {
			fmt.Println("Unmarshal error: ", err)
		}

		resp = append(resp, tmp)
	}
	for _, value := range resp {
		lib.Log("secretScan.log", "Repo:\t", value.SourceMetadata.Data.Git.Repository)
		lib.Log("secretScan.log", "Commit:\t", value.SourceMetadata.Data.Git.Commit)
		lib.Log("secretScan.log", "File:\t", value.SourceMetadata.Data.Git.File)
		lib.Log("secretScan.log", "Line:\t", value.SourceMetadata.Data.Git.Line)
		lib.Log("secretScan.log", "Raw:\t", value.Raw)
		lib.Log("secretScan.log", "Verified:\t", value.Verified)
		lib.Log("secretScan.log", "Time:\t", value.SourceMetadata.Data.Git.Timestamp)
		lib.Log("secretScan.log", "Email:\t", value.SourceMetadata.Data.Git.Email)
		lib.Log("secretScan.log", "\n")
	}
}
