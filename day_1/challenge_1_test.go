package main_test

import (
	"adventofcode-2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

type Tuple struct {
	First int
	Last  int
}

func ChallengeParser(content []string) ([]Tuple, error) {
	numberTuple := []Tuple{}
	re := regexp.MustCompile(`\d+`)

	for _, line := range content {
		matches := re.FindAllString(string(line), -1)

		if len(matches) < 2 {
			return nil, fmt.Errorf("Only one match found on line")
		}

		first, err := strconv.Atoi(matches[0])
		if err != nil {
			return nil, err
		}

		second, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, err
		}

		numberTuple = append(numberTuple, Tuple{First: first, Last: second})
	}

	return numberTuple, nil
}

func TestParseChallengeFile(t *testing.T) {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFileLineByLine("example.txt", nil)

	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}

	numberTuples, err := ChallengeParser(content)
	if err != nil {
		t.Fatalf("ChallengeParser failed: %v", err)
	}

	if len(numberTuples) != 6 {
		t.Fatalf("Expected 6 tuples, got %d", len(numberTuples))
	}

	if first, last := numberTuples[0].First, numberTuples[0].Last; first != 3 || last != 4 {
		t.Fatalf("First tuple is incorrect: expected (3, 4), got (%d, %d)", first, last)
	}
}
