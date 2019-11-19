package main

import (
	"fmt"
)

func binarySearch(leaderScores []int, myScore int) int {

	lowIndex := 0
	highIndex := len(leaderScores) - 1

	median := (lowIndex + highIndex) / 2

	if myScore == leaderScores[median] {
		return median
	} else if myScore > leaderScores[median] {
		return 1
	} else {
		return 2
	}

	// for lowIndex <= highIndex {
	// 	median := (lowIndex + highIndex) / 2
	// }
}

// ladderRanking receives as input the current Leaders' scores, as well as my personal scores
// and calculates my ranking in the leaderboard.
func ladderRanking(currentLeaders []int, myScores []int) []int {
	// Finding the ranking of the currentLeaders
	currentLeadersRanking := make([]int, len(currentLeaders))
	// currentLeaders is sorted descending, first element has the highest ranking
	currentLeadersRanking[0] = 1
	for i := 1; i < len(currentLeadersRanking); i++ {
		if currentLeaders[i] == currentLeaders[i-1] {
			currentLeadersRanking[i] = currentLeadersRanking[i-1]
		} else {
			currentLeadersRanking[i] = currentLeadersRanking[i-1] + 1
		}
	}

	myScoresRanking := make([]int, len(myScores))

	for i, myScore := range myScores {
		highestScore := currentLeaders[0]
		lowestScore := currentLeaders[len(currentLeaders)-1]

		if myScore > highestScore {
			myScoresRanking[i] = 1
		} else if myScore < lowestScore {
			myScoresRanking[i] = currentLeadersRanking[len(currentLeadersRanking)-1] + 1
		} else {
			myScoreIndex := binarySearch(currentLeaders, myScore)
			myScoresRanking[i] = currentLeadersRanking[myScoreIndex]
		}
	}
	return myScoresRanking
}

func main() {
	leaders := []int{110, 110, 80, 60, 60, 30, 25}
	personalScores := []int{3, 25, 50, 80, 90, 110, 120}
	// personalScores := []int{3}

	personalRanking := ladderRanking(leaders, personalScores)
	fmt.Println("My ranking is:", personalRanking)
}
