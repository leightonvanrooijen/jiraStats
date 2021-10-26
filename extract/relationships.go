package extract

import (
	"github.com/jira/server/db"
)

func UserRelationship(ctx *db.JiraDB, user *db.User, sprints []*db.Sprint) {
	ctx.CreateRelationships(user, sprints)
}
