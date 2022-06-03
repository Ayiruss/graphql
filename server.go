package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Ayiruss/bookstore/graph"
	"github.com/Ayiruss/bookstore/graph/directives"
	"github.com/Ayiruss/bookstore/graph/generated"
	"github.com/Ayiruss/bookstore/graph/utils/service"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()
	router.Use(service.Authenticate)
	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
