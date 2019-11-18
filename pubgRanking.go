package main

import (
	"fmt"
)

// uniqueLeaderScores receives as input the current Leaders' scores,
// and calculates how many times each score appeares.
func uniqueLeaderScores(currentLeaders []int) map[int]int {
	// Find unique leaders scores
	uniqueScores := make(map[int]int)

	for _, v := range currentLeaders {
		_, exists := uniqueScores[v]
		if exists == true {
			uniqueScores[v]++
		} else {
			uniqueScores[v] = 1
		}
	}

	return uniqueScores
}

func binarySearch(pointer *[]int, number int) int {
	// Recursive function ??
	return 30
}

// ladderRanking receives as input the current Leaders' scores, as well as my personal scores
// and calculates my ranking in the leaderboard.
func ladderRanking(currentLeaders []int, myScores []int) []int {
	// The map must include the ranking, not the duplicate count!!! ERROR!!
	// Find unique leaders scores
	uniqueScores := uniqueLeaderScores(currentLeaders)

	myRanking := make([]int, len(myScores))

	for index, myScore := range myScores {
		// TODO: Replace with switch
		highestScore := currentLeaders[0]
		lowestScore := currentLeaders[len(currentLeaders)-1]
		if myScore >= highestScore {
			myRanking[index] = 1
		} else if myScore == lowestScore {
			myRanking[index] = uniqueScores[myScore]
		} else if myScore < lowestScore {
			myRanking[index] = uniqueScores[lowestScore] + 1
		} else {
			myRanking[index] = binarySearch(&currentLeaders, myScore)
		}

	}
	return myRanking
}

func main() {
	leaders := []int{110, 110, 80, 60, 60, 30, 25}
	// personalScores := []int{3, 25, 50, 80, 90, 110, 120}
	personalScores := []int{3}

	personalRanking := ladderRanking(leaders, personalScores)
	fmt.Println("My ranking is:", personalRanking)
}
