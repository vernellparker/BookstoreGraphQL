package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vernellparker/BookstoreGraphQL/handlers/graphqlHandlers"
	"log"
)

func main() {
	//Sets up gin to handle requests
	r := gin.Default()
	r.POST("/query", graphqlHandlers.GraphqlHandler())
	r.GET("/", graphqlHandlers.PlaygroundHandler())

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
