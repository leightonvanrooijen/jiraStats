package db

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// AddUser takes a user object and maps it to the database
func (ctx JiraDB) CreateUser(user *User) (*User, error) {
	var mysqlErr *mysql.MySQLError

	err := ctx.db.Create(&user)

	if errors.As(err.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		dbUser, _ := ctx.GetUserByJiraID(user.JiraId)
		fmt.Printf(" user ID: %d", dbUser.ID)
		return dbUser, nil
	} else if err.Error != nil {
		return nil, fmt.Errorf("could not create user: %s\nerror: %s", user.Name, err.Error)
	}

	return user, nil
}

// GetUser fetches a user by id
func (ctx JiraDB) GetUser(id uint) (*User, error) {
	var user User
	if err := ctx.db.Preload("Issues").Preload("Sprints").First(&user, id); err.Error != nil {
		return nil, fmt.Errorf("%s", err.Error)
	}
	return &user, nil
}

// GetUser fetches a user by id
func (ctx JiraDB) GetUserByJiraID(id string) (*User, error) {
	var user User
	if err := ctx.db.First(&user, "jira_id = ?", id); err.Error != nil {
		return nil, fmt.Errorf("could not find user with the jira id of %s\nerror: %s", id, err.Error)
	}
	return &user, nil
}

// GetAllUsers fetches a users
func (ctx JiraDB) GetAllUsers() ([]User, error) {
	var users []User
	if err := ctx.db.Preload("Issues").Preload("Sprints").Find(&users); err.Error != nil {
		return nil, err.Error
	}
	return users, nil
}

// CreateRelationships creates all nessarcy user relationships
func (ctx JiraDB) CreateRelationships(
	user *User,
	sprints []*Sprint,
) error {
	if len(sprints) == 0 {
		return nil
	}

	user.Sprints = append(user.Sprints, sprints...)

	if err := ctx.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&user); err != nil {
		return nil
	}
	return nil
}
