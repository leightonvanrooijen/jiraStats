package extract

type Resp struct {
	Issues []*Issue `json:"issues"`
}

type Issue struct {
	Key      string `json:"key"`
	Fields   Fields `json:"fields"`
	UserID   int
	SprintID int
}

type Fields struct {
	Name         string    `json:"summary"`
	CompleteDate string    `json:"statuscategorychangedate"`
	Points       float64   `json:"customfield_10005"`
	User         User      `json:"assignee"`
	Status       Status    `json:"status"`
	Sprint       []*Sprint `json:"customfield_10007"`
}

type User struct {
	ID     string `json:"accountId"`
	Name   string `json:"displayName"`
	Active bool   `json:"active"`
	Type   string `json:"accountType"`
	Issues []Issue
}

type Status struct {
	Name string `json:"name"`
}

type Sprint struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	State     string `json:"state"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
