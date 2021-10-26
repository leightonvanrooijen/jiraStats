package extract

import (
	"sync"

	"github.com/jira/server/db"
)

func UpdateAll(ctx *db.JiraDB) {
	dbUsers := UserController(ctx)
	var wg sync.WaitGroup
	for _, dbUser := range dbUsers {
		wg.Add(1)
		go func(dbUser *db.User) {
			IssueController(ctx, dbUser)
			wg.Done()
		}(dbUser)
	}
	wg.Wait()
}

// func UpdateAll(ctx *db.JiraDB) {
// 	dbUsers := UserController(ctx)
// 	for _, dbUser := range dbUsers {
// 		IssueController(ctx, dbUser)
// 	}
// }
