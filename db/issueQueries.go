package db

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// AddIssue takes a issue object and maps it to the database
func (ctx JiraDB) CreateIssue(issue *Issue) (*Issue, error) {
	var mysqlErr *mysql.MySQLError

	err := ctx.db.Create(&issue)

	if errors.As(err.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		dbIssue, _ := ctx.GetIssueByKey(issue.Key)
		return dbIssue, nil
	} else if err.Error != nil {
		return nil, fmt.Errorf("could not create issue: %s\nerror: %s", issue.Name, err.Error)
	}
	return issue, nil
}

// GetIssue fetches a issue by id
func (ctx JiraDB) GetIssue(id uint) (*Issue, error) {
	var issue Issue
	if err := ctx.db.Preload("User").Preload("Sprint").First(&issue, id); err.Error != nil {
		return nil, fmt.Errorf("%s", err.Error)
	}
	return &issue, nil
}

// GetUser fetches a user by id
func (ctx JiraDB) GetIssueByKey(key string) (*Issue, error) {
	var issue Issue
	if err := ctx.db.First(&issue, "issue_key = ?", key); err.Error != nil {
		return nil, fmt.Errorf("could not find issue with the key of %s\nerror: %s", key, err.Error)
	}
	return &issue, nil
}

// GetAllIssues fetches all issues
func (ctx JiraDB) GetAllIssues() ([]Issue, error) {
	var issues []Issue
	if err := ctx.db.Preload("User").Preload("Sprint").Find(&issues); err.Error != nil {
		return nil, err.Error
	}
	return issues, nil
}
