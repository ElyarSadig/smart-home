package handler

import (
	"net/http"

	"github.com/elyarsadig/smart-home-iot/internal/domain/service"
)

type RoomHandler struct {
	Service *service.RoomService
}

func NewRoom(service *service.RoomService) *RoomHandler {
	return &RoomHandler{Service: service}
}

func (h *RoomHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := r.URL.Query().Get("name")
	if name == "" {
		returnError(w, "name cannot be empty", http.StatusBadRequest)
		return
	}
	device, err := h.Service.GetByName(ctx, name)
	if err != nil {
		returnError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	returnSuccess(w, device)
}
