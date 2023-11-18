package main

import (
	"os"
	"path/filepath"
	"strings"
)

type WordCounter interface {
	countWords(content string) int
}

type FileReader interface {
	ReadFile(filepath string) (string, error)
}

type Counter struct{}

type Reader struct{}

type WordCountProcessor struct {
	Counter Counter
	Reader  Reader
}

func (wcp WordCountProcessor) processFile(path string) (int, error) {
	content, err := wcp.Reader.ReadFile(path)
	if err != nil {
		return 0, err
	}
	count := wcp.Counter.countWords(content)
	return count, nil
}

func (c Counter) countWords(content string) int {
	words := strings.Fields(content)
	return len(words)
}

func (r Reader) ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func readFolder(path string, wcp WordCountProcessor) (int, error) {
	var totalWords int

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Process files with supported extensions
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if ext == ".txt" || ext == ".docx" || ext == ".pdf" {
				wordCount, err := wcp.processFile(path)
				if err != nil {
					return err
				}
				totalWords += wordCount
			}
		}
		return nil
	})
	return totalWords, err
}

func main() {

}
