package echo

import (
	"fx-test/api/config"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type handler struct {
	lgr   *zap.Logger
	route string
}

func NewEchoHandler(lgr *zap.Logger, cfg *config.APIConfig) *handler {
	return &handler{
		lgr:   lgr,
		route: cfg.Route.Echo,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.lgr.Info("Echo Handler Run")
	if _, err := io.Copy(w, r.Body); err != nil {
		h.lgr.Error("Failed to handle request: ", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *handler) Pattern() string {
	return h.route
}
