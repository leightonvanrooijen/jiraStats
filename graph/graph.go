package graph

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jira/server/db"
)

// Reads and parses the schema from file
// Associates root resolver, checks for errors along the way
func parseSchema(path string, resolver interface{}) *graphql.Schema {
	bstr, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Couldn't get the Graphql schema file", err)
	}

	schemaString := string(bstr)
	parsedSchema, err := graphql.ParseSchema(
		schemaString,
		resolver,
	)

	if err != nil {
		log.Fatal("Couldn't parse the Graphql schema", err)
	}

	return parsedSchema
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Serves GraphQL Playground on root
// Serves GraphQL endpoint at /graphql
func ConnectGraphqQL(db *db.JiraDB) {
	playground := http.FileServer(http.Dir("graph/graphqlPlayground"))

	http.Handle("/", playground)
	http.Handle("/graphql", CorsMiddleware(&relay.Handler{
		Schema: parseSchema("./graph/schema.graphql", &RootResolver{db: db}),
	}))

	fmt.Println("serving on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
