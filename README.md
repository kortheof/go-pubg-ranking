# go-pubg-ranking

## Exercise

### Summary

Marios is playing PubG and wants to climb to the top of the leaderboard. To do that, he wants to keep track of his ranking. The game uses Dense Ranking, so its leaderboard works like this:

The player who has the highest score gets ranked number 1 on the leaderboard.
All the players that have equal score get the same rank number, while the next one gets the immediate following rank.
For example, suppose that we have 4 players playing the game and their scores are *100, 80, 80, 60*. These players will get ranked *1, 2, 2, 3* respectively. If Marios plays the game 3 times and gather *75, 105, 50* points, he will be ranked *3rd, 1st and 4th* after each game.

### Function Signature:

ladderRanking(currentLeaders []int, myScores []int) []int

where currentLeaders is a slice of integers in **decreasing order** that represent the current leaderboard scores, myScores is also a slice of integers that represent Marios' scores in **ascending order** and finally the function returns a slice of integers that represent Marios' actual rank in the leaderboard after each game.

## Solution

1. Master Branch
This solution first retrieves the ranking of the current leaderbord scores, and then compares each one of MyScores to the leaderboard scores, hence the time complexity is O(n + n^2).

2. Bonus1 Branch
Some unit tests are added to ensure the functionality of the function

3. Bonus2 Branch
The function is solved using binary search for each of MyScores, thus reducing the complexity to O(n + nlogn)

4. Bonus3 Branch
The function is re-written, to cover the case where both input slices are unordered.
