package main

import (
	"adventofcode-2024/utils"
	"fmt"
)

func main() {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFileLineByLine("day_1/input.txt", nil)

	if err != nil {
		fmt.Errorf("fileUtils.ReadFileContents failed: %v", err)
	}

	listItems, err := ChallengeParser(content)
	if err != nil {
		fmt.Errorf("ChallengeParser failed: %v", err)
	}

	orderedListItems, err := SolveChallenge(listItems)
	if err != nil {
		fmt.Errorf("SolveChallenge 1 failed: %v", err)
	}

	fmt.Println("DAY 1")
	fmt.Printf("Challenge 1: %d\n", SumOfDistances(orderedListItems))
	fmt.Printf("Challenge 2: %d\n", SumOfSimularityScores(orderedListItems))
}
