package main

import (
	"fmt"
)

func ladderRanking(currentLeaders []int, myScores []int) []int {
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

	fmt.Println(uniqueScores)

	// Check myScores against the unique leader scores
	// Each of myScores' ranking number is 1 + the number of items
	// ranked above it that are distinct with respect to the ranking order.

	myRanking := make([]int, len(myScores))
	for index, myScore := range myScores {
		counter := 0
		for uniqueScore := range uniqueScores {
			if uniqueScore > myScore {
				counter++
			}
		}
		myRanking[index] = counter + 1
	}

	// return []int{6, 5, 4, 2, 2, 1, 1}
	return myRanking
}

func main() {
	leaders := []int{110, 110, 80, 60, 60, 30, 25}
	personalScores := []int{3, 25, 50, 80, 90, 110, 120}

	personalRanking := ladderRanking(leaders, personalScores)
	fmt.Println("My ranking is:", personalRanking)
}
