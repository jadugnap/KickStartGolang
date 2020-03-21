package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var debug bool = false
var consoleInput bool = false
var scanner *bufio.Scanner = bufio.NewScanner(strings.NewReader(testcases))

var testcases string = `2
9 4
18 23 23 15 15 6 6 6 11
9 4
19 24 24 16 16 7 7 7 12`

type intSlice []int

func main() {
	TestCases := ScanSliceInt(1)[0]
	if debug {
		fmt.Println("Num of TestCases, T =", TestCases)
	}

	for t := 1; t < 1+TestCases; t++ {
		answer := SolveCaseInt()
		fmt.Printf("Case#%d: %d\n", t, answer)
	}
}

func SolveCaseInt() (answer int) {
	vars := ScanSliceInt(2)
	num, pick := vars[0], vars[1]
	ratings := intSlice(ScanSliceInt(num))
	set := SetTeamRatings(ratings, pick, num)

	return CalculateAnswers(ratings, pick, set)
}

// SetTeamRatings from ratings, pick, num => find all possible set of teams' targetRating
func SetTeamRatings(ratings intSlice, pick, num int) (set map[int]struct{}) {
	sort.Ints(ratings)
	targets := ratings[pick-1:]

	set = make(map[int]struct{})
	for _, key := range targets {
		set[key] = struct{}{}
	}

	if debug {
		fmt.Printf("\nnum = %d, pick = %d\n", num, pick)
		fmt.Println("ratings[] :", ratings)
		fmt.Println("targets[] :", targets)
		fmt.Println("set[] :", set)
	}

	return set
}

// CalculateAnswers from set of teams' targetRating => find all possible training hours needed
func CalculateAnswers(ratings intSlice, pick int, set map[int]struct{}) (answer int) {
	answer = -1
	answers := []int{}

	for key, _ := range set {
		// lastIdx := ratings.SliceLastIndex(key, ratings.SliceFirstIndex(key))
		lastIdx := ratings.SliceLastIndex(key, 0)
		team := ratings[lastIdx+1-pick : lastIdx+1]
		targetRating := team[len(team)-1]
		hours := []int{}

		currentAnswer := 0
		for _, currentRating := range team {
			hour := targetRating - currentRating
			hours = append(hours, hour)
			currentAnswer += hour
		}
		answers = append(answers, currentAnswer)
		if answer < 0 || answer > currentAnswer {
			answer = currentAnswer
		}
		if debug {
			fmt.Println("team[]", key, ":", team)
			fmt.Println("hours[]", key, ":", hours)
		}
	}
	if debug {
		fmt.Println("answers[] :", answers, "\n")
	}

	return answer
}

// func (slice intSlice) SliceFirstIndex(toMatch int) int {
// 	for idx, val := range slice {
// 		if val == toMatch {
// 			return idx
// 		}
// 	}
// 	return -1
// }

func (slice intSlice) SliceLastIndex(toMatch, fromIdx int) int {
	for idx, val := range slice[fromIdx:] {
		if val > toMatch {
			return idx - 1 + fromIdx
		}
	}
	return len(slice) - 1
}

func ScanSliceInt(len int) []int {
	if consoleInput {
		return SSIConsole(len)
	} else {
		return SSIString(len)
	}
}

func SSIConsole(len int) []int {
	var slice []int
	var elem int

	for n := 0; n < len; n++ {
		fmt.Scanf("%v", &elem)
		slice = append(slice, elem)
	}
	return slice
}

func SSIString(len int) []int {
	var slice []int
	var elem int

	scanner.Scan()
	text := scanner.Text()
	strSlice := strings.Split(text, " ")
	for _, str := range strSlice {
		elem, _ = strconv.Atoi(str)
		slice = append(slice, elem)
	}
	return slice
}
