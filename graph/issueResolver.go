package graph

import (
	"fmt"
)

// Gets sprint for a issue
func (i *IssueResolver) Sprint() *SprintResolver {
	sprint, _ := i.db.GetSprint(uint(i.issue.SprintID))
	return &SprintResolver{db: i.db, sprint: sprint}
}

// Gets all issues
func (r *RootResolver) Issues() ([]*IssueResolver, error) {
	issues, err := r.db.GetAllIssues()
	if err != nil {
		return nil, fmt.Errorf("error [%s]: Could not fetch issues", err)
	}

	issueRxs := make([]*IssueResolver, len(issues))
	for i := range issues {
		issueRxs[i] = &IssueResolver{
			db:    r.db,
			issue: &issues[i],
		}
	}
	return issueRxs, nil
}

// Gets the user object for UserResolvers
func (r RootResolver) Issue(args struct{ ID int32 }) (*IssueResolver, error) {
	issue, err := r.db.GetIssue(uint(args.ID))
	if err != nil {
		return nil, fmt.Errorf("error [%s]: Could not find issue with Id: %d", err, args.ID)
	}

	return &IssueResolver{db: r.db, issue: issue}, nil
}

// Gers the ID for the issue
func (i *IssueResolver) ID() *int32 {
	userID := i.issue.ID
	id := int32(userID)

	return &id
}

// Gets name for the issue
func (i *IssueResolver) Name() *string {
	return &i.issue.Name
}

// Gers the key for the issue
func (i *IssueResolver) Key() *string {
	return &i.issue.Key
}

// Gets story points from the issue
func (i *IssueResolver) Points() *int32 {
	points := i.issue.Points
	point32 := int32(points)

	return &point32
}

// Gets staus of the issue
func (i *IssueResolver) Status() *string {
	return &i.issue.Status
}

// Gets complete date for the issue
func (i *IssueResolver) CompleteDate() *string {
	return &i.issue.CompleteDate
}

// Gets user addigned to the issue
func (i *IssueResolver) UserID() *int32 {
	userID := i.issue.UserID
	id := int32(userID)

	return &id
}

// Gets the ID of the sprint the issue was in
func (i *IssueResolver) SprintID() *int32 {
	sprintID := i.issue.SprintID
	id := int32(sprintID)

	return &id
}
