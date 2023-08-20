package main

import (
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"go.uber.org/zap"

	"ForkMatch/internal/routes"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled code.
	Version = "dev"

	// id is the hostname of the machine.
	id, _ = os.Hostname()
)

func main() {
	now := time.Now()

	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	logger := zap.Must(zap.NewProduction())
	logger.Info(
		"Starting ForkMatch",
		zap.String("version", Version),
		zap.String("id", id),
		zap.String("environment", string(config.Environment)),
	)

	pingRoute := routes.NewPingHandler(logger)
	http.Handle(pingRoute.Pattern(), pingRoute)

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)

	logger.Info(
		"Stopped ForkMatch",
		zap.String("version", Version),
		zap.String("id", id),
		zap.Duration("uptime", time.Since(now)),
	)
}
