package handler

import (
	"errors"
	"net/http"

	"github.com/elyarsadig/smart-home-iot/config"
	"github.com/elyarsadig/smart-home-iot/internal/domain/service"
)

type DeviceHandler struct {
	Service *service.DeviceService
}

func NewDevice(service *service.DeviceService) *DeviceHandler {
	return &DeviceHandler{Service: service}
}

func (h *DeviceHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := r.URL.Query().Get("name")
	if name == "" {
		returnError(w, "name cannot be empty", http.StatusBadRequest)
		return
	}
	device, err := h.Service.GetByName(ctx, name)
	if err != nil {
		if errors.Is(err, config.NotFoundError) {
			returnError(w, "device not found", http.StatusNotFound)
			return
		}
		returnError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	returnSuccess(w, device)
}
