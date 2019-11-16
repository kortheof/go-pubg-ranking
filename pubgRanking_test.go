package main

import "testing"

type ioTest struct {
	id                           int
	inCurrentLeaders, inMyScores []int
	outMyRanking                 []int
}

var mytest = []ioTest{
	ioTest{
		id: 1,
		inCurrentLeaders: []int{110, 110, 80, 60, 60, 30, 25},
		inMyScores: []int{3, 25, 50, 80, 90, 110, 120},
		outMyRanking: []int{6,5,4,2,2,1,1},
	},
}
func TestLadderRanking(t *testing.T) {
	for _,v := range mytest {
		out := ladderRanking(v.inCurrentLeaders, v.inMyScores)
		equal := true
		for i:= range out {
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