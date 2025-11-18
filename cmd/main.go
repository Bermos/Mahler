package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"time"

	v1 "github.com/Bermos/Platform/internal/api/v1"
	"github.com/Bermos/Platform/internal/app"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/spf13/cobra"
)

//go:embed all:../web/dist
var embeddedFrontend embed.FS

var version = "dev" // Set via ldflags during build

type Options struct {
	Debug bool   `doc:"Enable debug logging"`
	Host  string `doc:"Hostname to listen on."`
	Port  int    `doc:"Port to listen on." short:"p" default:"8080"`
}

func main() {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Mahler Platform", version))

	a := app.NewApp()
	v1.Register(api, a)

	// Serve embedded or filesystem frontend
	frontendHandler := serveFrontend()
	mux.Handle("/", frontendHandler)

	// Then, create the CLI.
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		slog.Info("Starting Mahler Platform",
			"version", version,
			"port", opts.Port,
			"host", opts.Host,
			"debug", opts.Debug)

		// Create the HTTP server.
		server := http.Server{
			Addr:         fmt.Sprintf("%s:%d", opts.Host, opts.Port),
			Handler:      mux,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		}

		hooks.OnStart(func() {
			slog.Info("Server listening", "addr", server.Addr)
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				slog.Error("Failed to start server", "error", err)
			}
		})

		hooks.OnStop(func() {
			slog.Info("Shutting down server...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				slog.Error("Server shutdown failed", "error", err)
			}
			slog.Info("Server stopped")
		})
	})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "openapi",
		Short: "Print the OpenAPI spec",
		Run: func(cmd *cobra.Command, args []string) {
			b, err := api.OpenAPI().YAML()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(b))
		},
	})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Mahler Platform %s\n", version)
		},
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}

// serveFrontend serves the Vue frontend from embedded files or filesystem
func serveFrontend() http.Handler {
	// Check if running in development mode (web/dist exists on filesystem)
	if _, err := os.Stat("web/dist"); err == nil {
		slog.Info("Serving frontend from filesystem (development mode)")
		return http.FileServer(http.Dir("web/dist"))
	}

	// Use embedded frontend
	slog.Info("Serving embedded frontend")
	frontendFS, err := fs.Sub(embeddedFrontend, "web/dist")
	if err != nil {
		slog.Error("Failed to load embedded frontend", "error", err)
		// Return a handler that serves a simple error page
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Frontend not available", http.StatusInternalServerError)
		})
	}

	return http.FileServer(http.FS(frontendFS))
}
