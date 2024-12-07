package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"
)

func ReadFile(filename string, processContent func([]byte) error) ([]byte, error) {
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

func ReadFileLineByLine(filename string, processLine func(string) error) ([]string, error) {
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

func TestReadFile(t *testing.T) {
	content, err := ReadFile("../example.txt", nil)
	if err != nil {
		t.Fatalf("ReadFileContents failed: %v", err)
	}

	var tmpSlices []string

	// We can do this straight away? (see: ReadFileLineByLine)
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		line := scanner.Text()
		// append is more efficient
		tmpSlices = slices.Insert(tmpSlices, len(tmpSlices), line)
	}

	if len(tmpSlices) != 6 {
		t.Fatal("Too few slices found.")
	}
}

func TestReadFileWithCallback(t *testing.T) {
	callback := func(content []byte) error {
		fmt.Printf("Processing bytes: %d\n", len(content))
		if len(content) == 0 {
			return fmt.Errorf("file content is empty") // Return error instead of using t.Fatal
		}

		var tmpSlices []string

		// Process content line by line
		scanner := bufio.NewScanner(strings.NewReader(string(content)))
		for scanner.Scan() {
			line := scanner.Text()
			// append is more efficient
			tmpSlices = append(tmpSlices, line)
		}

		// Check if the number of lines is as expected
		if len(tmpSlices) != 6 {
			return fmt.Errorf("expected 6 lines, found %d", len(tmpSlices)) // Return error
		}

		return nil
	}

	fmt.Println("callback:")
	_, err := ReadFile("../example.txt", callback)
	if err != nil {
		t.Fatalf("ReadFileContents failed: %v", err)
	}
}

func TestReadFileLineByLine(t *testing.T) {
	content, err := ReadFileLineByLine("../example.txt", nil)

	if err != nil {
		t.Fatalf("ReadFileContents failed: %v", err)
	}

	if len(content) != 6 {
		t.Fatal("Too few lines found.")
	}
}

func TestReadFileLineByLineWithCallabck(t *testing.T) {
	callback := func(line string) error {
		fmt.Printf("Processing line: %s\n", line)
		return nil
	}

	fmt.Println("callback")
	content, err := ReadFileLineByLine("../example.txt", callback)

	if err != nil {
		t.Fatalf("ReadFileContents failed: %v", err)
	}

	if len(content) != 6 {
		t.Fatal("Too few lines found.")
	}
}
