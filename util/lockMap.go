// This file contains the structs and methods for the maps with lock features (challenge, duel)
// and miscellanous structs (gamestate, problem, submission)

package util

import "sync"

type ChallengeMap struct {
	RWMutex sync.RWMutex
	Map     map[string][]string
}

type DuelMap struct {
	RWMutex sync.RWMutex
	Map     map[string]GameState
}

type GameState struct {
	Duelists []string
	Handles  []string
	Problems []Problem
}

type Problem struct {
	ID        string
	Solver    int
	Timestamp float64
}

type Submission struct {
	ID        string
	Timestamp float64
	Succesful bool
}
