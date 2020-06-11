package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/vernellparker/BookstoreGraphQL/handlers/graphqlHandlers"
	"github.com/vernellparker/BookstoreGraphQL/middleware"
	"log"
)

func main() {
	//Sets up gin to handle requests. Gin is also a 3rd party that allows for better HTTP, middleware, logging etc an
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, err := middleware.JwtHandler()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/", graphqlHandlers.PlaygroundHandler())

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})


	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/query", graphqlHandlers.GraphqlHandler())
	}

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
