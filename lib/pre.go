package lib

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadLines(filename string) []string {
	fd, err := os.Open(filename)
	var data []string
	if err != nil {
		log.Fatal(err)

	}
	defer fd.Close()
	buff := bufio.NewReader(fd)
	for {
		line, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		data = append(data, string(line))
	}
	return data
}

// Parse the text, return a slice

func Usage() {
	programName := os.Args[0]
	fmt.Println("[-] Press \"" + programName + " -h\" to get help.")
}

func Help() {
	programName := os.Args[0]
	fmt.Println("[+]", programName, "-file", "to to specify input file")
	fmt.Println("[+]", programName, "-ssh", "to detect ssh ports")
	fmt.Println("[+]", programName, "-https", "to check tls certificate")
	fmt.Println("[+]", programName, "-key", "to scan leaked keys in github repo")
}
