package graph

import (
	"fmt"
)

// Gets all issues for a user
func (s *SprintResolver) Issues() ([]*IssueResolver, error) {
	issues := s.sprint.Issues

	issueRxs := make([]*IssueResolver, len(issues))
	for i := range issues {
		issueRxs[i] = &IssueResolver{
			db:    s.db,
			issue: issues[i],
		}
	}
	return issueRxs, nil
}

// Gets the sprint object for Sprint
func (r RootResolver) Sprint(args struct{ ID int32 }) (*SprintResolver, error) {
	sprint, err := r.db.GetSprint(uint(args.ID))
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &SprintResolver{db: r.db, sprint: sprint}, nil
}

// Gets all sprints
func (r *RootResolver) Sprints() ([]*SprintResolver, error) {
	sprints, err := r.db.GetAllSprints()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	sprintRxs := make([]*SprintResolver, len(sprints))
	for i := range sprints {
		sprintRxs[i] = &SprintResolver{
			db:     r.db,
			sprint: &sprints[i],
		}
	}
	return sprintRxs, nil
}

// Gers the ID for the sprint
func (s *SprintResolver) ID() *int32 {
	sprintID := s.sprint.ID
	id := int32(sprintID)

	return &id
}

// Gers the jira ID for sprint
func (s *SprintResolver) JiraId() *int32 {
	sprintID := s.sprint.JiraId
	id := int32(sprintID)

	return &id
}

// Gets sprint name
func (s *SprintResolver) Name() *string {
	return &s.sprint.Name
}

// Gets state from sprint
func (s *SprintResolver) State() *string {
	return &s.sprint.State
}

// Gets start date for the sprint
func (s *SprintResolver) StartDate() *string {
	return &s.sprint.StartDate
}

// Gets end date for the sprint
func (s *SprintResolver) EndDate() *string {
	return &s.sprint.EndDate
}
