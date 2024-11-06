package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
	"net/url"
	"github.com/ferneediaz/gopher-socials/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/ferneediaz/gopher-socials/docs"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr    string
	apiURL  string
	db      dbConfig
	env     string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		
		// Correct Swagger configuration using apiURL
		docsURL := fmt.Sprintf("%s/v1/swagger/doc.json", app.config.apiURL)
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(docsURL), // Corrected docsURL
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
		))

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postsContextMiddleware)
				r.Get("/", app.getPostHandler)
				r.Delete("/", app.deletePostHandler)
				r.Patch("/", app.updatePostHandler)
			})
		})
		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.userContextMiddleware)
				r.Get("/", app.getUserHandler)
				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})
			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getUserFeedHandler)
			})
		})
	})
	return r
}

func (app *application) run(mux http.Handler) error {
	parsedURL, err := url.Parse(app.config.apiURL)
    if err != nil {
        log.Fatalf("Invalid EXTERNAL_URL: %v", err)
    }

	//Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = parsedURL.Host
	docs.SwaggerInfo.BasePath = "/v1"

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}

