package routes

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type PingHandler struct {
	log *zap.Logger
}

func NewPingHandler(log *zap.Logger) *PingHandler {
	return &PingHandler{log: log}
}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "Pong"); err != nil {
		p.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (p *PingHandler) Pattern() string {
	return "/ping"
}
