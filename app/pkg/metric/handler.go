package metric

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
}

// Register TODO исправить зависимость от httprouter
func (h *Handler) Register(router *httprouter.Router) { // Регистрация - принимает на себя роутер и регистрирует его
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}

// Heartbeat
// @Summary Heartbeat metric
// @Tags Metrics
// @Success 204
// @Failure 400
// @Router /api/heartbeat [get]
func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) { // Стук сердца - можно управлять состоянием приложения
	w.WriteHeader(204)
}
