package graph

import "github.com/jira/server/db"

type RootResolver struct {
	db *db.JiraDB
}

type UserResolver struct {
	db   *db.JiraDB
	user *db.User
}

type SprintResolver struct {
	db     *db.JiraDB
	sprint *db.Sprint
}

type IssueResolver struct {
	db    *db.JiraDB
	issue *db.Issue
}
