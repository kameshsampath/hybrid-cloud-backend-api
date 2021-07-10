package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/kameshsampath/hybrid-cloud-backend/docs"
	"github.com/kameshsampath/hybrid-cloud-backend/pkg/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	router *gin.Engine
)

// @title Hybrid Cloud Demo Backend API
// @version 1.0
// @description The backend API that processes the message from frontend

// @contact.name Kamesh Sampath
// @contact.email kamesh.sampath@solo.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1/api
// @query.collection.format multi
// @schemes http https
func main() {
	if mode := os.Getenv("GIN_MODE"); mode != "" {
		gin.SetMode(mode)
	}
	router = gin.Default()
	// this is liberal CORS settings only for demo
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	addRoutes()
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listent: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server shutting down server ...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced shutdown", err)
	}

	log.Println("Server Exiting")
}

func addRoutes() {
	endpoints := routes.NewEndpoints()

	v1 := router.Group("/v1/api")
	{
		//Health Endpoints accessible via /v1/api/health
		health := v1.Group("/health")
		{
			health.GET("/live", endpoints.Live)
			health.GET("/ready", endpoints.Ready)
		}
		//Backend endpoints
		backend := v1.Group("/process")
		{
			backend.POST("/", endpoints.Process)
		}
	}

	// the default path to get swagger json is :8080/swagger/docs.json
	// TODO enable/disable based on ENV variable
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
