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
	"github.com/sahilq312/workly/middleware"
	"github.com/sahilq312/workly/routes"
)

func init() {
	initializer.LoadEnvVariale()
	initializer.ConnectPostgresDatabase()
}

func main() {
	r := gin.New()

	// Set up health check routes
	setupRoutes(r)

	// Create a new HTTP server
	srv := &http.Server{
		Addr:    ":" + getPort(),
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go startServer(srv)

	// Wait for interrupt signal
	<-ctx.Done()
	gracefulShutdown(srv)
}

// setupRoutes initializes the server routes
func setupRoutes(r *gin.Engine) {
	r.GET("/", welcomeHandler)
	r.GET("/health", middleware.RequireAuth, healthCheckHandler)

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
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": "Server is healthy", "user": user})
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

func gracefulShutdown(srv *http.Server) {
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
