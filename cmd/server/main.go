package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang_rest_app/config"
	"golang_rest_app/controllers"
	db "golang_rest_app/db/sqlc"
	"golang_rest_app/routes"
	"log"
	"net/http"
)

var (
	server         *gin.Engine
	dbConn         *db.Queries
	AuthController controllers.AuthController
	PostController controllers.PostController
	AuthRoutes     routes.AuthRoutes
	PostRoutes     routes.PostRoutes
	ctx            context.Context
)

func init() {
	ctx = context.TODO()
	baseConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	conn, err := sql.Open(baseConfig.PostgresDriver, baseConfig.PostgresSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}
	dbConn = db.New(conn)

	AuthController = *controllers.NewAuthController(dbConn)
	PostController = *controllers.NewPostController(ctx, dbConn)

	AuthRoutes = routes.NewAuthRoutes(AuthController)
	PostRoutes = routes.NewRoutePost(PostController)
	server = gin.Default()
}

func main() {
	baseConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{baseConfig.Origin}
	corsConfig.AllowCredentials = true

	router := server.Group("/api")
	AuthRoutes.AuthRoute(router)
	PostRoutes.PostRoute(router)
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + baseConfig.ServerPort))
}
