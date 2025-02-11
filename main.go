package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"blankProject/application"
	"blankProject/common"
	"blankProject/server"
)

func main() {
	log.Info("Retrieving configuration")
	cfg, err := common.GetConfig()
	if nil != err {
		panic(fmt.Errorf("cannot get env variables: %w", err))
	}

	app := application.App{
		Config: cfg,
	}

	log.WithField("Port", app.Config.Port).Info("Starting server")
	panic(http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Port), server.NewServer()))
}
