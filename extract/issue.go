package extract

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jira/server/db"
	"github.com/jira/server/request"
)

func IssueController(ctx *db.JiraDB, dbUser *db.User) {
	issues := GetIssues(dbUser)
	dbSprints, updatedIssues := SprintController(ctx, issues, dbUser)
	SaveIssues(ctx, updatedIssues)
	UserRelationship(ctx, dbUser, dbSprints)
}

func GetIssues(dbuser *db.User) []*Issue {
	page := 0
	count := 1000
	var issues []*Issue
	for count >= 1000 {
		resp := request.GetIssues("https://ezyvet.atlassian.net/rest/api/3/search", dbuser.JiraId, page)

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var data Resp
		err = json.Unmarshal([]byte(responseData), &data)
		if err != nil {
			fmt.Println(err)
		}

		for _, issue := range issues {
			issue.UserID = dbuser.ID
		}

		issues = append(issues, data.Issues...)
		count = len(data.Issues)
		page++
	}

	return issues
}

func SaveIssues(ctx *db.JiraDB, issues []*Issue) []*db.Issue {
	var dbIssues []*db.Issue
	for _, issue := range issues {
		dbIssue := MapIssue(ctx, issue)
		svaedDbIssue, err := ctx.CreateIssue(dbIssue)
		if err != nil {
			fmt.Println(err)
		}

		if svaedDbIssue != nil {
			dbIssues = append(dbIssues, svaedDbIssue)
		}
	}
	return dbIssues
}

func MapIssue(ctx *db.JiraDB, issue *Issue) *db.Issue {
	return &db.Issue{
		Key:          issue.Key,
		Points:       int(issue.Fields.Points),
		Status:       issue.Fields.Status.Name,
		Name:         issue.Fields.Name,
		CompleteDate: issue.Fields.CompleteDate,
		UserID:       issue.UserID,
		SprintID:     issue.SprintID,
	}
}
