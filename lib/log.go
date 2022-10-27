package lib

// to write logs without write-conflicts
import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type SyncWriter struct {
	m      sync.Mutex
	Writer io.Writer
}

func (w *SyncWriter) Write(b []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()
	return w.Writer.Write(b)
}

// Use mutex to avoid write-conflict

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func Log(filename string, content any) {
	var f *os.File
	var err error
	if checkFileIsExist(filename) {
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
		// If the file exists, then open it
	} else {
		f, err = os.Create(filename)
		// If the file doesn't exist, then create one
	}
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	wr := &SyncWriter{sync.Mutex{}, f}
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Fprintln(wr, content)
	wg.Done()
	wg.Wait()
}
