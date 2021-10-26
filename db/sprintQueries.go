package db

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// AddSprint takes a sprint object and maps it to the database
func (ctx JiraDB) CreateSprint(sprint *Sprint) (*Sprint, error) {
	var mysqlErr *mysql.MySQLError

	resp := ctx.db.Create(&sprint)

	if errors.As(resp.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		dbSprint, _ := ctx.GetSprintByJiraID(sprint.JiraId)
		return dbSprint, nil
	} else if resp.Error != nil {
		return nil, fmt.Errorf("could not create sprint: %s\nerror: %s", sprint.Name, resp.Error)
	}
	return sprint, nil
}

// GetUser fetches a user by id
func (ctx JiraDB) GetSprintByJiraID(id int) (*Sprint, error) {
	var sprint Sprint
	if err := ctx.db.First(&sprint, "jira_id = ?", id); err.Error != nil {
		return nil, fmt.Errorf("could not find sprint with the JiraId of %d\nerror: %s", id, err.Error)
	}
	return &sprint, nil
}

// GetUser fetches a user by id
func (ctx JiraDB) GetSprint(id uint) (*Sprint, error) {
	var sprint Sprint
	if err := ctx.db.Preload("Issues").First(&sprint, id); err.Error != nil {
		return nil, fmt.Errorf("could not find sprint with the ID of %d\nerror: %s", id, err.Error)
	}
	return &sprint, nil
}

// GetAllUsers fetches a users
func (ctx JiraDB) GetAllSprints() ([]Sprint, error) {
	var sprints []Sprint
	if err := ctx.db.Preload("Issues").Find(&sprints); err.Error != nil {
		return nil, err.Error
	}
	return sprints, nil
}
