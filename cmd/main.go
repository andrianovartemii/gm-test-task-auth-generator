package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gm-test-task-auth-generator/internal/cmd"
	"os"
	"strconv"
)

func main() {
	host := ":8080"
	lifetime, err := strconv.Atoi(os.Getenv("AUTH_TOKEN_LIFETIME"))
	if err != nil {
		log.Panic("token lifetime must be number")
	}
	app, err := cmd.NewApp(host, lifetime)
	if err != nil {
		log.Panicf("error while create app, message %e", err)
	}
	if err == app.Run() {
		log.Panicf("error while running app, message %e", err)
	}
	defer func() {
		if err == app.Stop(context.Background()) {
			log.Panicf("error while running app, message %e", err)
		}
	}()
}
