package day_1

import (
	"adventofcode-2024/utils"
	"fmt"
)

func challenge_1() {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFileLineByLine("example.txt", nil)

	if err != nil {
		fmt.Errorf("fileUtils.ReadFileContents failed: %v", err)
	}

	listItems, err := ChallengeParser(content)
	if err != nil {
		fmt.Errorf("ChallengeParser failed: %v", err)
	}

	orderedListItems, err := OrderListemsAscending(listItems)
	if err != nil {
		fmt.Errorf("OrderListemsAscending failed: %v", err)
	}

	// Test orde results
	if orderedListItems[0].Left != 1 || orderedListItems[0].Right != 3 || orderedListItems[0].Distance != 2 {
		fmt.Errorf("Row 0: Expected (1, 3, 2), got (%d, %d, %d)", orderedListItems[0].Left, orderedListItems[0].Right, orderedListItems[0].Distance)
	}

	if orderedListItems[1].Left != 2 || orderedListItems[1].Right != 3 || orderedListItems[1].Distance != 1 {
		fmt.Errorf("Row 1: Expected (2, 3, 1), got (%d, %d, %d)", orderedListItems[1].Left, orderedListItems[1].Right, orderedListItems[1].Distance)
	}

	if orderedListItems[2].Left != 3 || orderedListItems[2].Right != 3 || orderedListItems[2].Distance != 0 {
		fmt.Errorf("Row 2: Expected (3, 3, 0), got (%d, %d, %d)", orderedListItems[2].Left, orderedListItems[2].Right, orderedListItems[2].Distance)
	}

	if orderedListItems[3].Left != 3 || orderedListItems[3].Right != 4 || orderedListItems[3].Distance != 1 {
		fmt.Errorf("Row 3: Expected (3, 4, 1), got (%d, %d, %d)", orderedListItems[3].Left, orderedListItems[3].Right, orderedListItems[3].Distance)
	}

	if orderedListItems[4].Left != 3 || orderedListItems[4].Right != 5 || orderedListItems[4].Distance != 2 {
		fmt.Errorf("Row 4: Expected (3, 5, 2), got (%d, %d, %d)", orderedListItems[4].Left, orderedListItems[4].Right, orderedListItems[4].Distance)
	}

	if orderedListItems[5].Left != 4 || orderedListItems[5].Right != 9 || orderedListItems[5].Distance != 5 {
		fmt.Errorf("Row 5: Expected (4, 9, 5), got (%d, %d, %d)", orderedListItems[5].Left, orderedListItems[5].Right, orderedListItems[5].Distance)
	}

	// Sum of weights (why is  there no map, reduce or sum?!?!)
	if SumOfDistances(orderedListItems) != 11 {
		fmt.Errorf("Expeced (11), got (%d)", SumOfDistances(orderedListItems))
	}
}
