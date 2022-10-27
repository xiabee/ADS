package secretScan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func Trufflehog() {
	url := "https://github.com/xiabee/security-test.git"
	// target repository

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

	cmd := exec.Command("trufflehog", "git", url, "--json")
	// To check by json mode

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // standard output
	cmd.Stderr = &stderr // standard error output
	err := cmd.Run()
	if err != nil {
		log.Fatal("Execution error:  ", err)
	}
	if stderr.Len() != 0 {
		log.Fatal("Stderr: ", stderr.Bytes())
	}
	// Error output
	// fmt.Println(string(stdout.Bytes()))

	var resp RESP
	err = json.Unmarshal(stdout.Bytes(), &resp)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
	}
	fmt.Println(resp.SourceMetadata.Data.Git.Commit)
	fmt.Println(resp.SourceMetadata.Data.Git.File)
}
