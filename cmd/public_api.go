package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/ngtrdai197/go-rabbitmq/domain/user"
	"github.com/ngtrdai197/go-rabbitmq/middleware"
)

func init() {
	rootCmd.AddCommand(publicCmd)
}

var publicCmd = &cobra.Command{
	Use:   "public-api",
	Short: "Serve public api application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, _ []string) {
		setup()
	},
}

func setup() {
	r := gin.New()
	r.Use(
		gin.Recovery(),
		middleware.GinLogger(),
	) // Add middleware here

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Register routes
	user.RegisterRoutes(r)

	// Setup gin server with graceful shutdown
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	// Start the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msgf("ListenAndServe: %v", err)
		}
	}()
	log.Info().Msg("Server started successfully")

	// Wait for an interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutting down server...")

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server shutdown failed: %v", err)
	}

	log.Info().Msg("Server shutdown successfully")
}
