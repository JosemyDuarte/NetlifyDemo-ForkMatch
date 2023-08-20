package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ForkMatch/internal/routes"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version = "dev"
	// flagconf is the config flag.
	//flagconf string

	// id is the hostname of the machine.
	id, _ = os.Hostname()
)

func main() {
	logger := zap.Must(zap.NewProduction())

	logger.Info(
		"Starting ForkMatch",
		zap.String("version", Version),
		zap.String("id", id),
	)

	r := gin.Default()
	pingRoute := routes.NewPingHandler(logger)
	r.GET(pingRoute.Pattern(), func(c *gin.Context) {
		pingRoute.ServeHTTP(c.Writer, c.Request)
	})

	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
