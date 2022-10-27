package lib

import (
	"bufio"
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
