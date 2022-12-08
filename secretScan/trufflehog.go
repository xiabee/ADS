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
	var truffleWg sync.WaitGroup
	urlist := lib.ReadLines(filename)
	truffleWg.Add(len(urlist))
	for i := range urlist {
		url := urlist[i]
		go func() {
			Trufflehog(url)
			truffleWg.Done()
		}()
	}
	truffleWg.Wait()
}

// use goroutine

func Trufflehog(url string) {
	outputFile := "secretScan.log"
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
		lib.Log(outputFile, "Repo:\t", value.SourceMetadata.Data.Git.Repository)
		lib.Log(outputFile, "Commit:\t", value.SourceMetadata.Data.Git.Commit)
		lib.Log(outputFile, "File:\t", value.SourceMetadata.Data.Git.File)
		lib.Log(outputFile, "Line:\t", value.SourceMetadata.Data.Git.Line)
		lib.Log(outputFile, "Raw:\t", value.Raw)
		lib.Log(outputFile, "Verified:\t", value.Verified)
		lib.Log(outputFile, "Time:\t", value.SourceMetadata.Data.Git.Timestamp)
		lib.Log(outputFile, "Email:\t", value.SourceMetadata.Data.Git.Email)
		lib.Log(outputFile, "\n")
	}
}
