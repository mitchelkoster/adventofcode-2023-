package main_test

import (
	"adventofcode-2024/utils"
	"fmt"
	"math"
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

		first, err := strconv.Atoi(matches[0])
		if err != nil {
			return nil, err
		}

		last, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, err
		}

		numberListItem = append(numberListItem, ListItem{
			Left:     first,
			Right:    last,
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
			Distance: int(math.Abs(float64(leftList[i]) - float64(rightList[i]))),
		})
	}

	return newListItems, nil
}

func SumOfDistances(listItems []ListItem) int {
	sum := 0
	for _, item := range listItems {
		sum += item.Distance
	}

	return sum
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

func TestChallengeSolution(t *testing.T) {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFileLineByLine("example.txt", nil)

	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}

	listItems, err := ChallengeParser(content)
	if err != nil {
		t.Fatalf("ChallengeParser failed: %v", err)
	}

	orderedListItems, err := OrderListemsAscending(listItems)
	if err != nil {
		t.Fatalf("OrderListemsAscending failed: %v", err)
	}

	// Test orde results
	if orderedListItems[0].Left != 1 || orderedListItems[0].Right != 3 || orderedListItems[0].Distance != 2 {
		t.Fatalf("Row 0: Expected (1, 3, 2), got (%d, %d, %d)", orderedListItems[0].Left, orderedListItems[0].Right, orderedListItems[0].Distance)
	}

	if orderedListItems[1].Left != 2 || orderedListItems[1].Right != 3 || orderedListItems[1].Distance != 1 {
		t.Fatalf("Row 1: Expected (2, 3, 1), got (%d, %d, %d)", orderedListItems[1].Left, orderedListItems[1].Right, orderedListItems[1].Distance)
	}

	if orderedListItems[2].Left != 3 || orderedListItems[2].Right != 3 || orderedListItems[2].Distance != 0 {
		t.Fatalf("Row 2: Expected (3, 3, 0), got (%d, %d, %d)", orderedListItems[2].Left, orderedListItems[2].Right, orderedListItems[2].Distance)
	}

	if orderedListItems[3].Left != 3 || orderedListItems[3].Right != 4 || orderedListItems[3].Distance != 1 {
		t.Fatalf("Row 3: Expected (3, 4, 1), got (%d, %d, %d)", orderedListItems[3].Left, orderedListItems[3].Right, orderedListItems[3].Distance)
	}

	if orderedListItems[4].Left != 3 || orderedListItems[4].Right != 5 || orderedListItems[4].Distance != 2 {
		t.Fatalf("Row 4: Expected (3, 5, 2), got (%d, %d, %d)", orderedListItems[4].Left, orderedListItems[4].Right, orderedListItems[4].Distance)
	}

	if orderedListItems[5].Left != 4 || orderedListItems[5].Right != 9 || orderedListItems[5].Distance != 5 {
		t.Fatalf("Row 5: Expected (4, 9, 5), got (%d, %d, %d)", orderedListItems[5].Left, orderedListItems[5].Right, orderedListItems[5].Distance)
	}

	// Sum of weights (why is  there no map, reduce or sum?!?!)
	if SumOfDistances(orderedListItems) != 11 {
		t.Fatalf("Expeced (11), got (%d)", SumOfDistances(orderedListItems))
	}
}
