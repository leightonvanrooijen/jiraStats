package db

import "gorm.io/gorm"

type Issue struct {
	gorm.Model
	Name         string
	Key          string `gorm:"uniqueIndex;size:255"`
	Points       int
	Status       string // make this a enum
	CompleteDate string // time.Time
	UserID       int
	User         User
	SprintID     int
	Sprint       Sprint
}

type Sprint struct {
	gorm.Model
	JiraId    int `gorm:"uniqueIndex;size:255"`
	Name      string
	State     string  // make this enum
	StartDate string  // time.Time
	EndDate   string  // time.Time
	Users     []*User `gorm:"many2many:users_sprints;"`
	Issues    []*Issue
}

type User struct {
	gorm.Model
	JiraId  string `gorm:"uniqueIndex;size:255"`
	Name    string
	Active  bool
	Type    string
	Issues  []*Issue
	Sprints []*Sprint `gorm:"many2many:users_sprints;"`
}
