package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
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
	r := gin.Default()

	// Set up custom CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow your frontend origin here
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	// Set up routes and start the server
	setupRoutes(r)

	srv := &http.Server{
		Addr:    ":" + getPort(),
		Handler: r,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go startServer(srv)

	<-ctx.Done()
	gracefulShutdown(srv)
}

// setupRoutes initializes the server routes
func setupRoutes(r *gin.Engine) {
	r.GET("/", welcomeHandler)
	r.GET("/health", middleware.RequireAuth, healthCheckHandler)
	r.GET("/company-health", middleware.CompanyAuth, healthCompanyCheckHandler)
	routes.AuthRoutes(r)
	routes.PostRoutes(r)
	routes.CompanyRoutes(r)
	routes.UserRoutes(r)
	routes.JobRoutes(r)
	routes.LikeRoutes(r)
	routes.CommentRoutes(r)
}

func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Workly"})
}

func healthCheckHandler(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": "Server is healthy", "user": user})
}

func healthCompanyCheckHandler(c *gin.Context) {
	company, _ := c.Get("company")
	c.JSON(http.StatusOK, gin.H{"message": "Company auth is Working Fine", "company": company})
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

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
