package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"go.uber.org/zap"

	"ForkMatch/internal/routes"
)

// App is the application for the service.
type App interface {
	// Run runs the application.
	Run() error
}

// NewApp returns a new application object.
func NewApp(config *Config) App {
	switch config.Environment {
	case EnvironmentLocal:
		return NewLocalApp(config)
	case EnvironmentAWS:
		return NewAWSApp(config)
	default:
		panic("unknown environment")
	}
}

// AWSApp is the application for the AWS environment.
type AWSApp struct {
	// Config is the configuration for the service.
	config *Config
}

func NewAWSApp(config *Config) *AWSApp {
	return &AWSApp{config: config}
}

func (app *AWSApp) Run() error {
	now := time.Now()

	logger := zap.Must(zap.NewProduction()).
		With(
			zap.String("id", id),
			zap.String("version", Version),
			zap.String("environment", string(app.config.Environment)),
		)
	logger.Info("Starting ForkMatch")

	pingRoute := routes.NewPingHandler(logger)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Hello, world!")
		if err != nil {
			logger.Error("Failed to write response", zap.Error(err))
		}
	})
	http.HandleFunc(pingRoute.Pattern(), pingRoute.ServeHTTP)

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)

	logger.Info(
		"Stopped ForkMatch",
		zap.Duration("uptime", time.Since(now)),
	)

	return nil
}

// LocalApp is the application for the local environment.
type LocalApp struct {
	// Config is the configuration for the service.
	config *Config
}

func NewLocalApp(config *Config) *LocalApp {
	return &LocalApp{config: config}
}

func (app *LocalApp) Run() error {
	now := time.Now()

	logger := zap.Must(zap.NewDevelopment()).
		With(
			zap.String("id", id),
			zap.String("version", Version),
			zap.String("environment", string(app.config.Environment)),
		)
	logger.Info("Starting ForkMatch")

	pingRoute := routes.NewPingHandler(logger)
	http.Handle(pingRoute.Pattern(), pingRoute)

	logger.Info("Listening on port " + app.config.ServingPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", app.config.ServingPort), nil); err != nil {
		logger.Error(
			"Failed to start ForkMatch",
			zap.Error(err),
		)

		return err
	}

	logger.Info(
		"Stopped ForkMatch",
		zap.Duration("uptime", time.Since(now)),
	)

	return nil
}
