package extract

import (
	"github.com/jira/server/db"
)

func SprintController(ctx *db.JiraDB, issues []*Issue, user *db.User) ([]*db.Sprint, []*Issue) {
	savedSprints, issues := GetSprints(ctx, issues, user)
	return savedSprints, issues
}

func GetSprints(ctx *db.JiraDB, issues []*Issue, user *db.User) ([]*db.Sprint, []*Issue) {
	var savedSprints []*db.Sprint
	var updatedIssues []*Issue
	for _, issue := range issues {
		if len(issue.Fields.Sprint) == 0 {
			continue
		}
		if issue.Fields.Sprint[0].Name == "" {
			continue
		}
		issue.UserID = user.ID
		savedSprint := SaveSprint(ctx, issue.Fields.Sprint[0])
		issue.SprintID = 0
		if savedSprint != nil {
			issue.SprintID = savedSprint.ID
			savedSprints = append(savedSprints, savedSprint)
		}
		
		updatedIssues = append(updatedIssues, issue)	
	}

	return savedSprints, updatedIssues
}

func SaveSprint(ctx *db.JiraDB, sprint *Sprint) *db.Sprint {
	dbSprint := MapSprint(sprint)
	savedSprint, _ := ctx.CreateSprint(dbSprint)
	return savedSprint
}

func MapSprint(sprint *Sprint) *db.Sprint {
	return &db.Sprint{
		JiraId:    sprint.ID,
		Name:      sprint.Name,
		State:     sprint.State,
		StartDate: sprint.StartDate,
		EndDate:   sprint.EndDate,
	}
}
