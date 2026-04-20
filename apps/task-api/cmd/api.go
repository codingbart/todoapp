package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/codingbart/todoapp/task-api/docs"
	"github.com/codingbart/todoapp/task-api/internal/config"
	db "github.com/codingbart/todoapp/task-api/internal/db/postgresql"
	"github.com/codingbart/todoapp/task-api/internal/health"
	"github.com/codingbart/todoapp/task-api/internal/logger"
	"github.com/codingbart/todoapp/task-api/internal/middleware"
	"github.com/codingbart/todoapp/task-api/internal/task"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Application interface {
	Mount() http.Handler
	Run(h http.Handler) error
}

type app struct {
	config  config.Config
	logger  logger.Logger
	queries *db.Queries
}

func NewApplication(config config.Config, logger logger.Logger, queries *db.Queries) Application {
	return &app{
		config:  config,
		logger:  logger,
		queries: queries,
	}
}

func (app *app) Mount() http.Handler {
	r := chi.NewRouter()

	app.configureSwagger()

	auth, err := middleware.NewAuthMiddleware(app.config.KeycloakJWKSURL, app.queries)
	if err != nil {
		app.logger.Error("failed to create auth middleware", "err", err)
	}

	healthService := health.NewService()
	healthHandler := health.NewHandler(healthService)

	taskService := task.NewService(app.queries)
	taskHandler := task.NewHandler(taskService)

	r.Route(app.config.BasePath, func(r chi.Router) {
		r.With(auth.Protect).Get("/health", healthHandler.GetHealthStatus)

		r.Group(func(r chi.Router) {
			r.Use(auth.Protect)
			r.Get("/users/{userId}/tasks", taskHandler.GetAll)
			r.Post("/users/{userId}/tasks", taskHandler.Create)
			r.Get("/users/{userId}/tasks/{id}", taskHandler.GetByID)
			r.Put("/users/{userId}/tasks/{id}", taskHandler.Update)
			r.Delete("/users/{userId}/tasks/{id}", taskHandler.Delete)
			r.Get("/users/{userId}/dashboard", taskHandler.GetDashboard)
		})
	})

	r.Handle("/swagger/*", app.swaggerHandler())

	return r
}

func (app *app) Run(h http.Handler) error {
	address := fmt.Sprintf("%s:%d", app.config.Host, app.config.Port)

	server := &http.Server{
		Addr:    address,
		Handler: h,
	}

	app.logger.Info("server started", "addr", address)

	return server.ListenAndServe()
}

func (app *app) configureSwagger() {
	docs.SwaggerInfo.BasePath = app.config.BasePath
	docs.SwaggerInfo.SwaggerTemplate = strings.NewReplacer(
		"KEYCLOAK_AUTH_URL", app.config.KeycloakAuthURL,
		"KEYCLOAK_TOKEN_URL", app.config.KeycloakTokenURL,
	).Replace(docs.SwaggerInfo.SwaggerTemplate)
}

func (app *app) swaggerHandler() http.Handler {
	return httpSwagger.Handler(
		httpSwagger.PersistAuthorization(true),
		httpSwagger.AfterScript(
			fmt.Sprintf(
				`window.ui.initOAuth({ clientId: %q, usePkceWithAuthorizationCodeGrant: true });`,
				app.config.KeycloakSwaggerClientID,
			),
		),
	)
}
