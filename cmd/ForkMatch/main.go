package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ForkMatch/internal/routes"
)

func main() {
	r := gin.Default()

	pingRoute := routes.NewPingHandler(zap.Must(zap.NewProduction()))
	r.GET(pingRoute.Pattern(), func(c *gin.Context) {
		pingRoute.ServeHTTP(c.Writer, c.Request)
	})

	err := r.Run(":80") // listen and serve on 0.0.0.0:80
	if err != nil {
		panic(err)
	}
}
