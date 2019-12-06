package githubclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func CreateIssue(title, body string, assignees []string, milestone int, labels []string) (bool, error) {
	reqParams := CreateIssueRequestParams{Title: title, Body: body}
	if milestone >= 0 {
		reqParams.Milestone = milestone
	}

	// TODO : validation

	reqParamJson, err := json.Marshal(reqParams)

	fmt.Printf("issue=%s\n", string(reqParamJson))
	req, err := http.NewRequest(
		"POST",
		IssuesURL,
		bytes.NewBuffer(reqParamJson),
	)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", " token "+AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return false, errors.New(fmt.Sprintf("create issue request failed: %d\n", resp.StatusCode))
	}

	return true, err
}

func ListIssuesForRepository() ([]Issue, error) {
	req, err := http.NewRequest("GET", IssuesURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("list issues for repository request failed: %d\n", resp.StatusCode))
	}

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func editIssue(issueNumber int, title, body string, assignees []string, milestone int, labels []string) (bool, error) {
	reqParam := EditIssueRequestParams{Title: title, Body: body}
	if milestone >= 0 {
		reqParam.Milestone = milestone
	}

	// TODO : validation

	reqParamJson, err := json.Marshal(reqParam)

	req, err := http.NewRequest("PATCH", IssuesURL+string(issueNumber), bytes.NewBuffer(reqParamJson))

	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", " token "+AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, errors.New(fmt.Sprintf("edit issue request failed: %d\n", resp.StatusCode))
	}

	return true, err
}
