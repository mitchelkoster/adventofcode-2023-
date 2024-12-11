package main_test

import (
	"adventofcode-2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

type ListItem struct {
	Left  int
	Right int
}

func ChallengeParser(content []string) ([]ListItem, error) {
	numberListItem := []ListItem{}
	re := regexp.MustCompile(`\d+`)

	for _, line := range content {
		matches := re.FindAllString(string(line), -1)

		if len(matches) < 2 {
			return nil, fmt.Errorf("Only one match found on line")
		}

		Left, err := strconv.Atoi(matches[0])
		if err != nil {
			return nil, err
		}

		second, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, err
		}

		numberListItem = append(numberListItem, ListItem{Left: Left, Right: second})
	}

	return numberListItem, nil
}

func TestParseChallengeFile(t *testing.T) {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFileLineByLine("example.txt", nil)

	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}

	listItems, err := ChallengeParser(content)
	if err != nil {
		t.Fatalf("ChallengeParser failed: %v", err)
	}

	if len(listItems) != 6 {
		t.Fatalf("Expected 6 ListItems, got %d", len(listItems))
	}

	if Left, Right := listItems[0].Left, listItems[0].Right; Left != 3 || Right != 4 {
		t.Fatalf("Left ListItem is incorrect: expected (3, 4), got (%d, %d)", Left, Right)
	}
}
