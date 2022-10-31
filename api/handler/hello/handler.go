package hello

import (
	"fmt"
	"fx-test/api/config"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type handler struct {
	lgr   *zap.Logger
	route string
}

func NewHelloHandler(lgr *zap.Logger, cfg *config.APIConfig) *handler {
	return &handler{
		lgr:   lgr,
		route: cfg.Route.Hello,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.lgr.Info("Hello Handler Run")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.lgr.Error("Failed to read request body: ", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if _, err = fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.lgr.Error("Failed to write response: ", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) Pattern() string {
	return h.route
}
