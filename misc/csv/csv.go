package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jira/server/misc/searchNested"
	"github.com/jira/server/request"
	"github.com/jira/server/utils"
)

// TODO this whole thing doesn't make sense

type Csv struct {
	writer *csv.Writer
}

func Create(name string) Csv {
	fileName := name + ".csv"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("message", err)
	}

	writer := csv.NewWriter(file)
	return Csv{writer: writer}
}

func (csv Csv) Write(data []string) {
	csv.writer.Write(data)
}

func (csv *Csv) Build(data map[string]interface{}) {
	var header []string
	header = append(header, "key", "name", "points", "status", "sprint")
	csv.Write(header)
	for _, issue := range data["issues"].([]interface{}) {
		key, ok := searchNested.Search(issue, "key")
		utils.CheckError(ok, "Couldn't find it")

		points, ok := searchNested.Search(issue, "customfield_10005")
		utils.CheckError(ok, "Couldn't find it")

		name, ok := searchNested.Search(issue, "summary")
		utils.CheckError(ok, "Couldn't find it")

		assignee, ok := searchNested.Search(issue, "assignee")
		utils.CheckError(ok, "Couldn't find it")

		assigneeName, ok := searchNested.Search(assignee, "displayName")
		utils.CheckError(ok, "Couldn't find it")

		sprint, ok := searchNested.Search(issue, "customfield_10007")
		utils.CheckError(ok, "Couldn't find it")

		sprintName, ok := searchNested.Search(sprint, "name")
		utils.CheckError(ok, "Couldn't find it")

		status, ok := searchNested.Search(issue, "status")
		utils.CheckError(ok, "Couldn't find it")

		statusName, ok := searchNested.Search(status, "name")
		utils.CheckError(ok, "Couldn't find it")

		fmt.Println("------Start-------")
		fmt.Println(name)
		fmt.Println(key)
		fmt.Println(points)
		fmt.Println(sprintName)
		fmt.Println(assigneeName)
		fmt.Println(statusName)
		// fmt.Println(sprint)
		fmt.Println("------End-------")

		var arr []string
		key = utils.ConvertString(key)
		name = utils.ConvertString(name)
		statusName = utils.ConvertString(statusName)
		sprintName = utils.ConvertString(sprintName)

		if points != nil {
			points = points.(float64)
			points = fmt.Sprintf("%f", points)
		} else {
			points = "0"
		}

		arr = append(arr, key.(string), name.(string), points.(string), statusName.(string), sprintName.(string))
		csv.Write(arr)
	}
}

func Build(user string) {
	resp := request.GetIssues("https://ezyvet.atlassian.net/rest/api/3/search", user, 1)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(responseData), &data)
	if err != nil {
		panic(err)
	}
	csv := Create("jira")
	csv.Build(data)
}
