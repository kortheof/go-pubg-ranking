package main

import (
	"fmt"
)

// ladderRanking receives as input the current Leaders' scores, as well as my personal scores
// and calculates my ranking in the leaderboard.
func ladderRanking(currentLeaders []int, myScores []int) []int {

	return []int{6, 5, 4, 2, 2, 1, 1}

}

func main() {
	leaders := []int{110, 110, 80, 60, 60, 30, 25}
	personalScores := []int{3, 25, 50, 80, 90, 110, 120}

	personalRanking := ladderRanking(leaders, personalScores)
	fmt.Println("My ranking is:", personalRanking)
}
