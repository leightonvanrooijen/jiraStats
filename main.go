package main

import (
	"github.com/jira/server/db"
	"github.com/jira/server/extract"
	"github.com/jira/server/graph"
)

func main() {
	db := db.ConnectDB()
	extract.UpdateAll(db)
	graph.ConnectGraphqQL(db)
}

// TODO create graphql server
// TODO dockerise
// TODO create mysql db
