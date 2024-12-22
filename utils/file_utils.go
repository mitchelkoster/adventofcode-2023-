package utils

import (
	"bufio"
	"os"
)

type FileUtils struct {
	Filename string
}

func (p *FileUtils) ReadFile(filename string, processContent func([]byte) error) ([]byte, error) {
	fh, err := os.ReadFile(filename) // auto closes?
	if err != nil {
		return nil, err
	}

	// Callback provided
	if processContent != nil {
		if err := processContent(fh); err != nil {
			return nil, err
		}
	}

	return fh, nil
}

func (p *FileUtils) ReadFileLineByLine(filename string, processLine func(string) error) ([]string, error) {
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fh.Close()

	// appending is more efficient than inserts used below
	var lines []string
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()

		// Callback provided
		if processLine != nil {
			if err := processLine(line); err != nil {
				return nil, err
			}
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
