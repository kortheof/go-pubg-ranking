package main

import (
	"fmt"
)

func binarySearch(pointer []int, number int) int {
	// Recursive function ??
	return 6
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
	// for index, myScore := range myScores {
	// 	// TODO: Replace with switch
	// 	highestScore := currentLeaders[0]
	// 	lowestScore := currentLeaders[len(currentLeaders)-1]
	// 	if myScore >= highestScore {
	// 		myRanking[index] = 1
	// 	} else if myScore == lowestScore {
	// 		myRanking[index] = uniqueScores[myScore]
	// 	} else if myScore < lowestScore {
	// 		myRanking[index] = uniqueScores[lowestScore] + 1
	// 	} else {
	// 		myRanking[index] = binarySearch(&currentLeaders, myScore)
	// 	}

	// }
	// return myRanking
}

func main() {
	leaders := []int{110, 110, 80, 60, 60, 30, 25}
	personalScores := []int{3, 25, 50, 80, 90, 110, 120}
	// personalScores := []int{3}

	personalRanking := ladderRanking(leaders, personalScores)
	fmt.Println("My ranking is:", personalRanking)
}
