package application

import (
	"os"

	"github.com/vandi37/Calculator/config"
	"github.com/vandi37/Calculator/internal/http/handler"
	"github.com/vandi37/Calculator/internal/http/server"
	"github.com/vandi37/Calculator/pkg/calc_service"
	"github.com/vandi37/Calculator/pkg/logger"
	"github.com/vandi37/vanerrors"
)

type Application struct {
	config string
}

func New(config string) *Application {
	return &Application{config}
}

func (a *Application) Run() {
	// Adding json errors mode
	vanerrors.DefaultLoggerOptions.ShowAsJson = true

	// Creating logger
	logger := logger.New(os.Stderr)

	// Loading config
	config, err := config.LoadConfig(a.config)
	if err != nil {
		logger.Fatalln(err)
	}

	// Crating calc service
	service := calc_service.New(logger)
	// Adding logging
	service.DoLog = config.DoLog

	// Creating handler
	handler := handler.NewHandler(config.Path, service)

	// Creating server
	server := server.New(handler, config.Port)

	// Running server
	err = server.Run()
	if err != nil {
		logger.Fatalln(err)
	}
	// The program end
}
