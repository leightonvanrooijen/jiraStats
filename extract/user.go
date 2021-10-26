package extract

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jira/server/db"
	"github.com/jira/server/request"
)

func UserController(db *db.JiraDB) []*db.User {
	users := GetUsers()
	dbUsers := SaveUsers(db, *users)

	return dbUsers
}

func GetUsers() *[]User {
	page := 0
	count := 1000
	var users []User
	for count >= 1000 {
		resp := request.GetUsers(page)

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var respUsers []User
		err = json.Unmarshal([]byte(responseData), &respUsers)
		if err != nil {
			panic(err)
		}

		users = append(users, respUsers...)

		count = len(respUsers)
		page++
	}

	return &users
}

func SaveUsers(ctx *db.JiraDB, users []User) []*db.User {
	var dbUsers []*db.User
	for _, user := range users {
		if !ValidUser(user) {
			continue
		}

		dbUser := MapUser(user)

		savedDbUser, err := ctx.CreateUser(dbUser)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if savedDbUser != nil {
			dbUsers = append(dbUsers, savedDbUser)
		}
	}
	return dbUsers
}

func MapUser(user User) *db.User {
	return &db.User{
		JiraId: user.ID,
		Name:   user.Name,
		Active: user.Active,
		Type:   user.Type,
	}
}

func ValidUser(user User) bool {
	if !user.Active {
		return false
	} else if user.Type != "atlassian" {
		return false
	}

	return true
}
