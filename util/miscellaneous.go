package util

import (
	"strings"

	"math/rand"
)

var codeforcesTag map[string]int = map[string]int{
	"implementation":            0,
	"math":                      0,
	"greedy":                    0,
	"dp":                        0,
	"data structures":           0,
	"brute force":               0,
	"constructive algorithms":   0,
	"graphs":                    0,
	"sortings":                  0,
	"binary search":             0,
	"dfs and similar":           0,
	"trees":                     0,
	"strings":                   0,
	"number theory":             0,
	"combinatorics":             0,
	"special":                   0,
	"geometry":                  0,
	"bitmasks":                  0,
	"two pointers":              0,
	"dsu":                       0,
	"shortest paths":            0,
	"probabilities":             0,
	"divide and conquer":        0,
	"hashing":                   0,
	"games":                     0,
	"flows":                     0,
	"interactive":               0,
	"matrices":                  0,
	"string suffix structures":  0,
	"fft":                       0,
	"graph matchings":           0,
	"ternary search":            0,
	"expression parsing":        0,
	"meet-in-the-middle":        0,
	"2-sat":                     0,
	"chinese remainder theorem": 0,
	"schedules":                 0,
}

func TagSubset(a []string) bool {
	for _, i := range a {
		if _, exists := codeforcesTag[i]; !exists {
			return false
		}
	}
	return true
}

func Parse(args []string) ([]string, []string, bool) {
	var includeTags, excludeTags []string
	for _, tag := range args {
		if tag[0] == '+' {
			// include
			includeTags = append(includeTags, strings.Replace(tag[1:], "-", " ", -1))
		} else if tag[0] == '~' {
			// exclude
			excludeTags = append(excludeTags, strings.Replace(tag[1:], "-", " ", -1))
		} else {
			return nil, nil, false
		}
	}
	return includeTags, excludeTags, true
}

func RandomSample[T any](s []T, cnt int) ([]T, bool) {
	if cnt > len(s) {
		return nil, false
	}

	indices := rand.Perm(len(s))

	var res []T
	for i := 0; i < cnt; i++ {
		res = append(res, s[indices[i]])
	}
	return res, true
}

// 1 is X, 2 is O
func IsGameOver(board []int) (bool, bool) {
	win := []bool{false, false}

	// rows
	for i := 0; i < 9; i += 3 {
		if board[i] == board[i+1] && board[i+1] == board[i+2] && board[i] != 0 {
			win[board[i]-1] = true
		}
	}

	// columns
	for i := 0; i < 3; i += 1 {
		if board[i] == board[i+3] && board[i+3] == board[i+6] && board[i] != 0 {
			win[board[i]-1] = true
		}
	}

	// diagonal
	if board[0] == board[4] && board[4] == board[8] && board[0] != 0 {
		win[board[0]-1] = true
	}

	if board[2] == board[4] && board[4] == board[6] && board[2] != 0 {
		win[board[2]-1] = true
	}
	return win[0], win[1]
}
