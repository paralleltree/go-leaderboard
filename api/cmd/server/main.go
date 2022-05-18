package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paralleltree/go-leaderboard/internal/config"
)

func main() {
	env := config.GetEnv()
	router := gin.New()
	server := http.Server{Addr: fmt.Sprintf(":%d", env.Port), Handler: router}

	done := make(chan struct{})
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Fprintf(os.Stderr, "Unexpected error: %v\n", err)
			}
		}
		close(done)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-done:
	case <-quit:
	}

	fmt.Fprintln(os.Stderr, "Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Server forced to shutdown: %v\n", err)
	}

	fmt.Fprintln(os.Stderr, "Exiting.")
}
