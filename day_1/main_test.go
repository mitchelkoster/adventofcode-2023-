package main

import (
	"adventofcode-2024/utils"
	"fmt"
	"testing"
)

func TestChallenge1Solution(t *testing.T) {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFileLineByLine("input.txt", nil)

	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}

	listItems, err := ChallengeParser(content)
	if err != nil {
		t.Fatalf("ChallengeParser failed: %v", err)
	}

	orderedListItems, err := SolveChallenge(listItems)
	if err != nil {
		t.Fatalf("SolveChallenge failed: %v", err)
	}

	fmt.Printf("Solution: %d\n", SumOfDistances(orderedListItems))
}

func TestChallenge2Solution(t *testing.T) {
	fileUtils := utils.FileUtils{}
	content, err := fileUtils.ReadFileLineByLine("input.txt", nil)

	if err != nil {
		t.Fatalf("fileUtils.ReadFileContents failed: %v", err)
	}

	listItems, err := ChallengeParser(content)
	if err != nil {
		t.Fatalf("ChallengeParser failed: %v", err)
	}

	orderedListItems, err := SolveChallenge(listItems)
	if err != nil {
		t.Fatalf("SolveChallenge failed: %v", err)
	}

	fmt.Printf("Solution: %d\n", SumOfSimularityScores(orderedListItems))
}
