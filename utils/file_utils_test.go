package utils_test

import (
	"adventofcode-2024/utils"
	"bufio"
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFile("day_1_example.txt", nil)

	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}

	var tmpSlices []string

	// We can do this straight away? (see: fileUtils.ReadFileLineByLine)
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
	fileUtils := utils.FileUtils{}

	callback := func(content []byte) error {
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

	_, err := fileUtils.ReadFile("day_1_example.txt", callback)
	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}
}

func TestReadFileLineByLine(t *testing.T) {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFileLineByLine("day_1_example.txt", nil)

	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}

	if len(content) != 6 {
		t.Fatal("Too few lines found.")
	}
}

func TestReadFileLineByLineWithCallabck(t *testing.T) {
	fileUtils := utils.FileUtils{}

	callback := func(line string) error {
		return nil
	}

	content, err := fileUtils.ReadFileLineByLine("day_1_example.txt", callback)

	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}

	if len(content) != 6 {
		t.Fatal("Too few lines found.")
	}
}
