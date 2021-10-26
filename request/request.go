package request

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
)

type ReqArgs struct {
	url  string
	body []byte
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func RedirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+BasicAuth("leighton.vanrooijen@ezyvet.com", ""))
	return nil
}

func GetIssues(url string, user string, page int) *http.Response {
	body := []byte(fmt.Sprintf(`{
		"jql":"assignee = %s AND type in (Bug, Data-fix, Task, Story)",
		"maxResults": 1000,
		"startAt": %d
	}`, user, page))

	args := ReqArgs{
		url:  url,
		body: body,
	}

	return Post(args)
}

func GetUsers(page int) *http.Response {
	args := ReqArgs{
		url: fmt.Sprintf("https://ezyvet.atlassian.net/rest/api/3/users/search?maxResults=1000&startAt=%d", page),
	}

	req, _ := Get(args)

	return req
}

func Get(args ReqArgs) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: RedirectPolicyFunc,
	}

	req, err := http.NewRequest("GET", args.url, nil)
	if err != nil {
		return nil, fmt.Errorf("error: GET request failed: %s", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Basic "+BasicAuth("leighton.vanrooijen@ezyvet.com", ""))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("GET request failed: %s\n", err)
	}

	return resp, nil
}

func Post(args ReqArgs) *http.Response {
	client := &http.Client{
		CheckRedirect: RedirectPolicyFunc,
	}

	req, err := http.NewRequest("POST", args.url, bytes.NewBuffer(args.body))
	if err != nil {
		fmt.Printf("POST request failed: %s\n", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Basic "+BasicAuth("leighton.vanrooijen@ezyvet.com", ""))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("POST request failed: %s\n", err)
	}

	return resp
}
