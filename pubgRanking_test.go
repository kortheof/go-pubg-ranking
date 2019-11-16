package main

import "testing"

type ioTest struct {
	id                           int
	inCurrentLeaders, inMyScores []int
	outMyRanking                 []int
}

var mytest = []ioTest{
	ioTest{
		id:               1,
		inCurrentLeaders: []int{110, 110, 80, 60, 60, 30, 25},
		inMyScores:       []int{3, 25, 50, 80, 90, 110, 120},
		outMyRanking:     []int{6, 5, 4, 2, 2, 1, 1},
	},
	ioTest{
		id:               2,
		inCurrentLeaders: []int{100, 80, 80, 60},
		inMyScores:       []int{50, 75, 105},
		outMyRanking:     []int{4, 3, 1},
	},
	ioTest{
		id:               3,
		inCurrentLeaders: []int{100, 80, 80, 60},
		inMyScores:       []int{102, 103},
		outMyRanking:     []int{1, 1},
	},
	ioTest{
		id:               4,
		inCurrentLeaders: []int{100, 80, 80, 60},
		inMyScores:       []int{20, 30},
		outMyRanking:     []int{4, 4},
	},
}

func TestLadderRanking(t *testing.T) {
	for _, v := range mytest {
		out := ladderRanking(v.inCurrentLeaders, v.inMyScores)

		equal := true
		for i := range out {
			if out[i] != v.outMyRanking[i] {
				equal = false
				break
			}
		}
		if equal == false {
			t.Errorf("Expected %v but received %v for test id %d", v.outMyRanking, out, v.id)
		}
	}
}
