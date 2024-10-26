// Package main is the entry point of the application
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/routes"
)

// init function runs before main, used for initialization
func init() {
	initializer.LoadEnvVariale()
	initializer.ConnectPostgresDatabase()
}

// main function is the entry point of the application
func main() {
	r := gin.New() // Create a new Gin router

	// Set up health check routes
	setupRoutes(r)

	// Create a new HTTP server
	srv := &http.Server{
		Addr:    ":" + getPort(), // Use a helper function to get the port
		Handler: r,
	}

	// Set up graceful shutdown context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start the server in a goroutine
	go startServer(srv)

	// Wait for interrupt signal
	<-ctx.Done()
	gracefulShutdown(srv)
}

// setupRoutes initializes the server routes
func setupRoutes(r *gin.Engine) {
	r.GET("/", welcomeHandler)
	r.GET("/health", healthCheckHandler)

	// Set up routes for authentication, posts, companies, users, jobs, likes, and comments
	routes.AuthRoutes(r)
	routes.PostRoutes(r)
	routes.CompanyRoutes(r)
	routes.UserRoutes(r)
	routes.JobRoutes(r)
	routes.LikeRoutes(r)
	routes.CommentRoutes(r)
}

// welcomeHandler handles requests to the root path
func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Workly"})
}

// healthCheckHandler handles health check requests
func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Server is healthy"})
}

// getPort retrieves the server port from the environment variable
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// startServer starts the HTTP server
func startServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

// gracefulShutdown attempts to shut down the server gracefully
func gracefulShutdown(srv *http.Server) {
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to shut down the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
