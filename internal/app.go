package internal

import (
	"context"
	"net/http"
	"time"

	"github.com/Andrmist/it-revolution-test-mine/internal/services"
	"github.com/Andrmist/it-revolution-test-mine/internal/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs" // docs is generated by Swag CLI, you have to import it.
)

func Run(ctx context.Context, serverCtx types.ServerContext) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(contextMiddleware(serverCtx))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://stats.m0e.space", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/api/transform", services.TransformLink)
	r.Get("/api/original/{id}", services.GetOriginalLink)
	r.Get("/api/statistics", services.GetStatistics)

	go func() {
		if err := http.ListenAndServe("0.0.0.0:8081", r); err != nil {
			serverCtx.Log.Fatal(err)
		}
	}()
}

func contextMiddleware(serverCtx types.ServerContext) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "server", serverCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
