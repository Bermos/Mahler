package main

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/Bermos/Platform/internal/api/v1"
	"github.com/Bermos/Platform/internal/app"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/spf13/cobra"
	"log/slog"
	"net/http"
	"time"
)

type Options struct {
	Debug bool   `doc:"Enable debug logging"`
	Host  string `doc:"Hostname to listen on."`
	Port  int    `doc:"Port to listen on." short:"p" default:"8080"`
}

func main() {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Platform", "1.0.0"))

	a := app.NewApp()
	v1.Register(api, a)

	// Then, create the CLI.
	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		fmt.Printf("I was run with debug:%v host:%v port%v\n",
			opts.Debug, opts.Host, opts.Port)

		// Create the HTTP server.
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", opts.Port),
			Handler: mux,
		}

		hooks.OnStart(func() {
			// Start your server here
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				slog.Error("Failed to start server", "error", err)
			}
		})

		hooks.OnStop(func() {
			// Gracefully shutdown your server here
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			server.Shutdown(ctx)
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

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
