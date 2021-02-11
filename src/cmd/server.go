package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/matthausen/gql-example/cmd/graph"
	"github.com/matthausen/gql-example/cmd/graph/generated"
	"github.com/matthausen/gql-example/cmd/graph/model"
	"github.com/matthausen/gql-example/pkg/postgres"
)

const defaultPort = "8080"

var port = os.Getenv("PORT")
var dbUserName = os.Getenv("POSTGRES_USER")
var dbPassword = os.Getenv("POSTGRES_PASSWORD")
var dbURL = os.Getenv("POSTGRES_URL")
var dbName = os.Getenv("POSTGRES_DB")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userService := &postgres.UserService{
		DbUserName: "graphql_example",
		DbPassword: "graphql_password",
		DbURL:      "localhost",
		DbName:     "graphql",
	}

	err := userService.Initialise()
	if err != nil {
		log.Fatal(err)
	}

	config := generated.Config{Resolvers: &graph.Resolver{User: userService}}
	config.Directives.HasRole = hasRoleDirective

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func hasRoleDirective(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (res interface{}, err error) {
	log.Printf("Inside hasRoleDirective - ignore the role check for now")
	return next(ctx)
}
