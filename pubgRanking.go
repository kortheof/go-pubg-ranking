package main

import (
	"fmt"
)

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
	for i, v := range myScores {
		for j := len(currentLeaders) - 1; j >= 0; j-- {
			if v < currentLeaders[j] {
				myScoresRanking[i] = currentLeadersRanking[j] + 1
				break
			} else if v == currentLeaders[j] {
				myScoresRanking[i] = currentLeadersRanking[j]
				break
			}
			myScoresRanking[i] = 1
		}
	}
	return myScoresRanking
}

func main() {
	leaders := []int{110, 110, 80, 60, 60, 30, 25}
	personalScores := []int{3, 25, 50, 80, 90, 110, 120}

	personalRanking := ladderRanking(leaders, personalScores)
	fmt.Println("My ranking is:", personalRanking)
}
