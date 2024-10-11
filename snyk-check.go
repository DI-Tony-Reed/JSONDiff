package main

import (
	"errors"
	"fmt"
	"strings"
)

// Original author: https://github.com/hezro/snyk-code-pr-diff/tree/main
// Went this route as the original code does not have a go.mod and cannot be
// pulled in with Go's package management
func check(baseline map[string]interface{}, feature map[string]interface{}) (string, error) {
	var results = strings.Builder{}

	// Extract the "results" array from the Baseline scan
	baselineResults, ok := extractResults(baseline)
	if !ok {
		return "", errors.New("failed to extract 'results' from the Baseline scan")
	}

	// Extract the "results" array from the PR scan
	prResults, ok := extractResults(feature)
	if !ok {
		return "", errors.New("failed to extract 'results' from the PR scan")
	}

	// Find the indices of new fingerprints from the PR results
	newIndices := findNewFingerprintIndices(baselineResults, prResults)

	// Extract the new issues objects from the PR results
	newIssues := extractNewIssues(prResults, newIndices)

	// Count the number of new issues found from the PR results
	//issueCount := len(newIssues)

	// Output the new issues from the PR results
	for _, result := range newIssues {
		level, message, uri, startLine := extractIssueData(result)
		level = strings.Replace(level, "note", "Low", 1)
		level = strings.Replace(level, "warning", "Medium", 1)
		level = strings.Replace(level, "error", "High", 1)
		results.WriteString(fmt.Sprintf("✗ Severity: [%s]\n", level))
		results.WriteString(fmt.Sprintf("Path: %s\n", uri))
		results.WriteString(fmt.Sprintf("Start Line: %d\n", startLine))
		results.WriteString(fmt.Sprintf("Message: %s\n", message))
		results.WriteString("\n")
	}

	//// Output the count new issues found from the PR results
	//if issueCount > 0 {
	//	results.WriteString("\n")
	//	results.WriteString(fmt.Sprintf("Total issues found: %d\n", issueCount))
	//
	//	// Replace the "results" array in the PR scan with only the new issues found
	//	feature["runs"].([]interface{})[0].(map[string]interface{})["results"] = newIssues
	//
	//	// Convert the new PR data to JSON
	//	updatedPRScan, err := json.Marshal(feature)
	//	if err != nil {
	//		return "", errors.New("failed to convert updated data to JSON")
	//	}
	//
	//	// Write the updated PR diff scan to a file
	//	//err = ioutil.WriteFile("snyk_code_pr_diff_scan.json", updatedPRScan, 0644)
	//	//if err != nil {
	//	//	log.Fatalf("Failed to write updated data to file: %v", err)
	//	//}
	//	//fmt.Printf("\n")
	//	//fmt.Println("Results saved in usnyk_code_pr_diff_scan.json")
	//}

	results.WriteString("\n")
	results.WriteString("No issues found!")

	return results.String(), nil

}

// Extract the "results" array from the JSON data
func extractResults(data map[string]interface{}) ([]interface{}, bool) {
	runs, ok := data["runs"].([]interface{})
	if !ok {
		return nil, false
	}

	if len(runs) > 0 {
		results, ok := runs[0].(map[string]interface{})["results"].([]interface{})
		if !ok {
			return nil, false
		}
		return results, true
	}

	return nil, false
}

// Find the indices of the new fingerprints in the PR results array
func findNewFingerprintIndices(baselineResults, prResults []interface{}) []int {
	var newIndices []int

	for i, prResult := range prResults {
		prObject := prResult.(map[string]interface{})
		if prFingerprints, ok := prObject["fingerprints"].(map[string]interface{}); ok {
			matchFound := false
			for _, baselineResult := range baselineResults {
				baselineObject := baselineResult.(map[string]interface{})
				if baselineFingerprints, ok := baselineObject["fingerprints"].(map[string]interface{}); ok {
					// Ignore the "identity" key
					delete(baselineFingerprints, "identity")
					delete(prFingerprints, "identity")

					match := fmt.Sprint(prFingerprints) == fmt.Sprint(baselineFingerprints)
					if match {
						matchFound = true
						break
					}
				}
			}
			if !matchFound {
				newIndices = append(newIndices, i)
			}
		}
	}

	return newIndices
}

// Extract new issues objects from the PR "results" array
func extractNewIssues(results []interface{}, indices []int) []interface{} {
	var newIssues []interface{}

	for _, idx := range indices {
		newIssues = append(newIssues, results[idx])
	}

	return newIssues
}

// Extract new issue data from the results to output to the console
func extractIssueData(result interface{}) (string, string, string, int) {
	resultObj := result.(map[string]interface{})
	level := resultObj["level"].(string)
	message := resultObj["message"].(map[string]interface{})["text"].(string)
	locations := resultObj["locations"].([]interface{})
	uri := locations[0].(map[string]interface{})["physicalLocation"].(map[string]interface{})["artifactLocation"].(map[string]interface{})["uri"].(string)
	startLine := locations[0].(map[string]interface{})["physicalLocation"].(map[string]interface{})["region"].(map[string]interface{})["startLine"].(float64)
	return level, message, uri, int(startLine)
}
