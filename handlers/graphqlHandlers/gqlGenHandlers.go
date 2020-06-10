package graphqlHandlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vernellparker/BookstoreGraphQL/graph"
	"github.com/vernellparker/BookstoreGraphQL/graph/generated"
)

//Defines the Graphql handler
func GraphqlHandler() gin.HandlerFunc{
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc  {
	h := playground.Handler("GraphQL", "/query")
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

