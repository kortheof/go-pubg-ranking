package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

var stackSize int = 0
var stack = new(node)

func push(v int) {
	if stack == nil {
		stack = &node{v, nil}
		stackSize = 1
		// return true
	}
	temp := &node{v, nil}
	temp.next = stack
	stack = temp
	stackSize++
	// return true
}

func pop() (int, bool) {
	if stackSize == 0 {
		return 0, false
	}
	if stackSize == 1 {
		stackSize = 0
		poppedValue := stack.value
		stack = nil
		return poppedValue, true
	}
	poppedValue := stack.value
	stack = stack.next
	stackSize--
	return poppedValue, true
}

// ladderRanking receives as input the current Leaders' scores, as well as my personal scores
// and calculates my ranking in the leaderboard.
func ladderRanking(currentLeaders []int, myScores []int) []int {
	// currentLeaders is sorted descending, add the scores to a stack
	for _, v := range currentLeaders {
		// if the score already exists in the stack, skip it
		// the stack will hold only unique scores
		// bottom of stack -> highest score
		// top of stack -> lowest score
		if v == stack.value {
			continue
		}
		push(v)
	}

	// compare myScores with the scores in the stack
	// myScores is sorted ascending, and the comparison begins with the top
	// of the stack which holds the lowest leaderboard score
	myRanking := make([]int, len(myScores))

	for i := 0; i <= len(myScores)-1; {
		switch {
		// The stack is empty, so my ranking is definitely 1
		case stack == nil:
			myRanking[i] = 1
			i++

		case myScores[i] < stack.value:
			myRanking[i] = stackSize + 1
			i++

		case myScores[i] == stack.value:
			myRanking[i] = stackSize
			i++

		case myScores[i] > stack.value:
			for myScores[i] > stack.value {
				_, hasItems := pop()
				print(hasItems)
				// if hasItems == false {
				if stackSize == 0 {
					myRanking[i] = 1
					i++
					break
				}
			}
		}
	}
	return myRanking
}

func main() {
	leaders := []int{110, 110, 80, 60, 60, 30, 25}
	personalScores := []int{3, 25, 50, 80, 90, 110, 120, 130}
	// personalScores := []int{3}

	personalRanking := ladderRanking(leaders, personalScores)
	fmt.Println("My ranking is:", personalRanking)
	fmt.Println(stackSize)
}
