package graph

import (
	"fmt"
)

// Gets all issues for a user
func (u *UserResolver) Issues() ([]*IssueResolver, error) {
	issues := u.user.Issues

	issueRxs := make([]*IssueResolver, len(issues))
	for i := range issues {
		issueRxs[i] = &IssueResolver{
			db:    u.db,
			issue: issues[i],
		}
	}
	return issueRxs, nil
}

// Gets all issues for a user
func (u *UserResolver) Sprints() ([]*SprintResolver, error) {
	sprints := u.user.Sprints
	fmt.Println(sprints)

	sprintRxs := make([]*SprintResolver, len(sprints))
	for i := range sprints {
		sprintRxs[i] = &SprintResolver{
			db:     u.db,
			sprint: sprints[i],
		}
	}
	return sprintRxs, nil
}

// Gets all users
func (r *RootResolver) Users() ([]*UserResolver, error) {
	users, err := r.db.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("error [%s]: Could not fetch users", err)
	}

	userRxs := make([]*UserResolver, len(users))
	for i := range users {
		userRxs[i] = &UserResolver{
			db:   r.db,
			user: &users[i],
		}
	}
	return userRxs, nil
}

// Gets the user object for UserResolvers
func (r RootResolver) User(args struct{ ID int32 }) (*UserResolver, error) {
	user, err := r.db.GetUser(uint(args.ID))
	if err != nil {
		return nil, fmt.Errorf("error [%s]: Could not find user with Id: %d", err, args.ID)
	}

	return &UserResolver{db: r.db, user: user}, nil
}

// Gers the ID for users
func (u *UserResolver) ID() *int32 {
	userID := u.user.ID
	id := int32(userID)

	return &id
}

// Gers the ID for users
func (u *UserResolver) JiraId() *string {
	return &u.user.JiraId
}

// Gets name for user
func (u *UserResolver) Name() *string {
	return &u.user.Name
}

// Gets name from user
func (u *UserResolver) Active() *bool {
	return &u.user.Active
}

// Gets name from user
func (u *UserResolver) Type() *string {
	return &u.user.Type
}
