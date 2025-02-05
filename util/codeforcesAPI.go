package util

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

const codeforcesURL string = "https://codeforces.com/api/"

// get all problems with the given criteria (rating, tags, specific problems to exclude)
func GetProblems(targetRating float64, includeTags []string, excludeTags []string, excludeProblems map[string]int) []Problem {
	resp, err := http.Get(codeforcesURL + "problemset.problems") // request http for all problems
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	problems := data["result"].(map[string]interface{})["problems"].([]interface{}) // parse to a map
	problemsFiltered := make(map[string]Problem)                                    // result map

	// filtering problems based on the given criteria, unique by problem name (may erase unnecessary duplicates)
	for _, p := range problems {
		pm := p.(map[string]interface{})

		// if already in list
		if _, exist := problemsFiltered[pm["name"].(string)]; exist {
			continue
		}

		// if targetRating is set and out of range
		rating, ok := pm["rating"].(float64)
		if targetRating != 0 && (!ok || rating < targetRating-300 || rating > targetRating+300) {
			continue
		}

		// if tags are set and not matching
		tags := pm["tags"].([]interface{})
		ok = true
		includeTagsCount := 0
		for _, t := range tags {
			if slices.Contains(excludeTags, t.(string)) {
				ok = false
				break
			}

			if slices.Contains(includeTags, t.(string)) {
				includeTagsCount++
			}
		}
		if !ok || includeTagsCount < len(includeTags) {
			continue
		}
		// fmt.Println(rating, tags)

		id := strconv.Itoa(int(pm["contestId"].(float64))) + "/" + pm["index"].(string)

		if _, exists := excludeProblems[id]; exists {
			continue
		}

		// add to map with name as key
		problemsFiltered[pm["name"].(string)] = Problem{
			ID:     id,
			Solver: 0,
		}
	}

	// return only the values
	return slices.Collect(maps.Values(problemsFiltered))
}

// get submissions that are attempted by the handle (submitted and not compilation error)
func GetAttemptedSubmissions(handle string) map[string]Submission {
	resp, err := http.Get(codeforcesURL + "user.status?handle=" + handle) // http request
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	submissions := data["result"].([]interface{})       // parse to map
	attemptedSubmissions := make(map[string]Submission) // result map

	// filter submissions
	for _, sub := range submissions {
		subm := sub.(map[string]interface{})
		problem := subm["problem"].(map[string]interface{})
		// skip if compilation error
		if subm["verdict"] != "COMPILATION_ERROR" {
			// fmt.Println(problem)
			contestId, ok := problem["contestId"]
			if !ok {
				// if not in a contest, i.e. no contestId, skip
				continue
			}
			// add to map with key as contestId + index
			key := strconv.Itoa(int(contestId.(float64))) + "/" + problem["index"].(string)
			attemptedSubmissions[key] = Submission{
				ID:        key,
				Timestamp: subm["creationTimeSeconds"].(float64),
				Succesful: true,
			}
		}
	}
	return attemptedSubmissions
}

// get submissions that are successful by the handle
func GetSuccesfulSubmissions(handle string) map[string]Submission {
	resp, err := http.Get(codeforcesURL + "user.status?handle=" + handle) // http request
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	submissions := data["result"].([]interface{})        // parse to map
	successfulSubmissions := make(map[string]Submission) // result map

	// loop through submissions
	for _, sub := range submissions {
		subm := sub.(map[string]interface{})
		problem := subm["problem"].(map[string]interface{})

		// only choose if the verdict is OK
		if subm["verdict"] == "OK" {
			// add to map with key as contestId + index
			key := strconv.Itoa(int(problem["contestId"].(float64))) + "/" + problem["index"].(string)
			successfulSubmissions[key] = Submission{
				ID:        key,
				Timestamp: subm["creationTimeSeconds"].(float64),
				Succesful: true,
			}
		}
	}
	return successfulSubmissions
}

// check if user exists
func UserExists(handles []string) bool {
	// http request
	resp, err := http.Get(codeforcesURL + "user.info?handles=" + strings.Join(handles, ";") + "&checkHistoricHandles=true")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	// status OK if user exists
	return data["status"] == "OK"
}
