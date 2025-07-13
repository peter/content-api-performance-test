package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/seenthis-ab/content-api/config"
	"github.com/seenthis-ab/content-api/handlers"
	"github.com/seenthis-ab/content-api/middleware"
	"github.com/seenthis-ab/content-api/models"
	"go.uber.org/zap"
)

// Options for the CLI. Pass `--port` or set the `SERVICE_PORT` env var.
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

// Use the shared interface and Content struct from models package

func main() {
	logger := config.GetLogger()
	defer config.CloseLogger()

	contentStore, err := models.GetContentStore()
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
		router := chi.NewMux()

		// Add logging middleware
		router.Use(middleware.LoggingMiddleware())

		api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

		// Initialize content handlers
		contentHandlers := handlers.NewContentHandlers(contentStore)

		// Register content endpoints
		huma.Post(api, "/content", contentHandlers.CreateContent)
		huma.Get(api, "/content/{id}", contentHandlers.GetContent)
		huma.Get(api, "/content", contentHandlers.ListContent)
		huma.Put(api, "/content/{id}", contentHandlers.UpdateContent)
		huma.Delete(api, "/content/{id}", contentHandlers.DeleteContent)

		// Tell the CLI how to start your router.
		hooks.OnStart(func() {
			// Configure HTTP server for high concurrency
			server := &http.Server{
				Addr:         ":" + strconv.Itoa(options.Port),
				Handler:      router,
				ReadTimeout:  30 * time.Second,
				WriteTimeout: 30 * time.Second,
				IdleTimeout:  120 * time.Second,
			}

			logger.Info("Starting server with high concurrency settings",
				zap.Int("port", options.Port),
			)
			if err := server.ListenAndServe(); err != nil {
				logger.Error("Error starting server", zap.Error(err))
			}
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
