package day_1

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
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
