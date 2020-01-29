package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/confus1on/meetmeup/ent"
	go_gqlgen "github.com/confus1on/meetmeup/graphql"
	"github.com/confus1on/meetmeup/postgres"
)

const defaultPort = "8080"

var client *ent.Client

func main() {
	client, err := postgres.Open()
	if err != nil {
		log.Fatalf("Error connected to PostgreSQL %+v ", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := go_gqlgen.Config{Resolvers: &go_gqlgen.Resolver{
		MeetupsRepo: postgres.MeetupsRepository{Client: client},
		UsersRepo:   postgres.UsersRepository{Client: client},
	}}

	queryHandler := handler.GraphQL(go_gqlgen.NewExecutableSchema(config))

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", go_gqlgen.DataLoaderMiddleware(context.Background(), client, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
