package cmd

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gm-test-task-auth-generator/internal/services"
	"gm-test-task-auth-generator/internal/tokenManager"
	handlers "gm-test-task-auth-generator/internal/transport/http"
	"net/http"
	"time"
)

type App struct {
	server *http.Server
}

func NewApp(host string, tokenLifetime int) (*App, error) {
	log.Infof("Generate secret")
	secret, err := tokenManager.GenerateRandomSecret(40)
	log.Infof("The token for tokenManager was generated successfully")
	if err != nil {
		return nil, err
	}
	manager := &tokenManager.Manager{
		Secret:   secret,
		Lifetime: time.Duration(tokenLifetime) * time.Minute,
	}
	handler := handlers.NewHttpHandler(services.NewAuthService(manager))
	log.Infof("Then app was initted successfully")
	return &App{
		server: &http.Server{
			Addr:    host,
			Handler: handler.Init(),
		},
	}, nil
}

func (app *App) Run() error {
	log.Infof("Start server on %s", app.server.Addr)
	log.Infof("Good luck!")
	return app.server.ListenAndServe()
}

func (app *App) Stop(ctx context.Context) error {
	log.Infof("Stop server on %s", app.server.Addr)
	return app.server.Shutdown(ctx)
}
