package main

import (
	"bufio"
	"fmt"
	"githubclient"
	"os"
	"strconv"
	"strings"
)

const (
	title = iota
	body
	milestone
)

var inputCheckStatus int

func main() {

	for {
		fmt.Printf(">")
		stdin := bufio.NewScanner(os.Stdin)
		if !stdin.Scan() {
			os.Exit(1)
		}
		input := stdin.Text()

		if len(input) == 0 {
			//fmt.Printf("input len=0\n")
			continue
		}

		cmd := strings.Split(input, " ")
		//fmt.Printf("cmd[0]=%s\n", cmd[0])
		switch cmd[0] {
		case "issues":
			if len(cmd) == 1 {
				showIssueList()
			} else if len(cmd) > 1 {
				arg := cmd[1]
				i, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Printf("invalid argument:%s\n", arg)
				} else {
					showIssue(i)
				}
			}
		case "create":
			createIssue()
		case "edit":
			editIssue()
		default:
			fmt.Printf("unrecognize command: %s\n", cmd[0])
		}
	}
}

func showIssueList() {
	issues, err := githubclient.ListIssuesForRepository()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	fmt.Printf("list issues for repository count=%d\n", len(issues))
	for _, issue := range issues {
		fmt.Printf("%v #%-5d %9.9s %.55s\n", issue.CreatedAt, issue.Number, issue.User.Login, issue.Title)
		fmt.Printf("\t%3s\n", issue.Body)
	}
}

func showIssue(num int) {
	issues, err := githubclient.ListIssuesForRepository()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	var issue *githubclient.Issue
	for _, item := range issues {
		if item.Number == num {
			issue = &item
			break
		}
	}

	fmt.Printf("number: #%-5d\ntitle: %.55s\nuser:%9.9s\ncreated at:%v\n", issue.Number, issue.Title, issue.User.Login, issue.CreatedAt)
	fmt.Printf("body:%s\n", issue.Body)
}

func createIssue() error {
	inputCheckStatus = title

	stdin := bufio.NewScanner(os.Stdin)
	var titleTxt string
	var bodyTxt string
	var milestoneNum int

	for {
		var message string
		switch inputCheckStatus {
		case title:
			message = "enter title"
		case body:
			message = "enter body"
		case milestone:
			message = "enter milestone (number)"
		}

		fmt.Printf("%s\n", message)

		if !stdin.Scan() {
			fmt.Printf("scan error\n")
			os.Exit(1)
		}

		input := stdin.Text()

		switch inputCheckStatus {
		case title:
			if input == "" {
				fmt.Printf("input error. title can't be empty.\n")
				continue
			}

			titleTxt = input
		case body:
			bodyTxt = input
		case milestone:
			if input == "" {
				milestoneNum = -1
			} else {
				num, err := strconv.Atoi(input)
				if err != nil {
					fmt.Printf("input error. %s is not a number.\n", input)
					continue
				}
				milestoneNum = num
			}
		}

		inputCheckStatus++
		if inputCheckStatus > milestone {
			break
		}
	}

	_, err := githubclient.CreateIssue(titleTxt, bodyTxt, nil, milestoneNum, nil)
	if err != nil {
		return err
	}

	return nil
}

func editIssue() {
	// TODO
}

func closeIssue() {
	// TODO
}
