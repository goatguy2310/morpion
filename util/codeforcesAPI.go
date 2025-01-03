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

func GetProblems(targetRating float64, includeTags []string, excludeTags []string, excludeProblems map[string]int) []Problem {
	resp, err := http.Get(codeforcesURL + "problemset.problems")
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

	problems := data["result"].(map[string]interface{})["problems"].([]interface{})

	problemsFiltered := make(map[string]Problem)
	for _, p := range problems {
		pm := p.(map[string]interface{})
		if _, exist := problemsFiltered[pm["name"].(string)]; exist {
			continue
		}

		rating, ok := pm["rating"].(float64)
		if targetRating != 0 && (!ok || rating < targetRating-300 || rating > targetRating+300) {
			continue
		}

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

		problemsFiltered[pm["name"].(string)] = Problem{
			ID:     id,
			Solver: 0,
		}
	}

	return slices.Collect(maps.Values(problemsFiltered))
}

func GetAttemptedSubmissions(handle string) map[string]Submission {
	resp, err := http.Get(codeforcesURL + "user.status?handle=" + handle)
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

	submissions := data["result"].([]interface{})
	attemptedSubmissions := make(map[string]Submission)

	for _, sub := range submissions {
		subm := sub.(map[string]interface{})
		problem := subm["problem"].(map[string]interface{})
		if subm["verdict"] != "COMPILATION_ERROR" {
			fmt.Println(problem)
			contestId, ok := problem["contestId"]
			if !ok {
				continue
			}
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

func GetSuccesfulSubmissions(handle string) map[string]Submission {
	resp, err := http.Get(codeforcesURL + "user.status?handle=" + handle)
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

	submissions := data["result"].([]interface{})
	successfulSubmissions := make(map[string]Submission)

	for _, sub := range submissions {
		subm := sub.(map[string]interface{})
		problem := subm["problem"].(map[string]interface{})
		if subm["verdict"] == "OK" {
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

func UserExists(handles []string) bool {
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
	return data["status"] == "OK"
}
