package main_test

import (
	"adventofcode-2024/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"testing"
)

type ListItem struct {
	Left     int
	Right    int
	Distance int
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

		numberListItem = append(numberListItem, ListItem{
			Left:     Left,
			Right:    second,
			Distance: 0,
		})
	}

	return numberListItem, nil
}

func OrderListemsAscending(listItems []ListItem) ([]ListItem, error) {
	leftList, rightList := []int{}, []int{}
	for _, item := range listItems {
		leftList = append(leftList, item.Left)
		rightList = append(rightList, item.Right)
	}

	if !sort.IntsAreSorted(leftList) {
		sort.Ints(leftList)
	}

	if !sort.IntsAreSorted(rightList) {
		sort.Ints(rightList)
	}

	newListItems := []ListItem{}
	for i, _ := range leftList {
		newListItems = append(newListItems, ListItem{
			Left:     leftList[i],
			Right:    rightList[i],
			Distance: 0, // TODO: Let's be lazy and take in distance
		})
	}

	return newListItems, nil
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

	if left, right := listItems[0].Left, listItems[0].Right; left != 3 || right != 4 {
		t.Fatalf("Left ListItem is incorrect: expected (3, 4), got (%d, %d)", left, right)
	}
}

func TestOrderAscendingByList(t *testing.T) {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFileLineByLine("example.txt", nil)

	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}

	listItems, err := ChallengeParser(content)
	if err != nil {
		t.Fatalf("ChallengeParser failed: %v", err)
	}

	for _, val := range listItems {
		fmt.Println(val)
	}
	fmt.Println("====")

	orderedListItems, err := OrderListemsAscending(listItems)
	if err != nil {
		t.Fatalf("OrderListemsAscending failed: %v", err)
	}

	// Test first row
	if left, right := orderedListItems[0].Left, orderedListItems[0].Right; left != 1 && right != 3 {
		t.Fatalf("Left ListItem is incorrect: expected (1, 3), got (%d, %d)", left, right)
	}

}
