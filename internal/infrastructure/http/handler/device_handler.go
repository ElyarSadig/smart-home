package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/elyarsadig/smart-home-iot/config"
	"github.com/elyarsadig/smart-home-iot/internal/domain/service"
)

type DeviceHandler struct {
	Service *service.DeviceService
}

func NewDevice(service *service.DeviceService) *DeviceHandler {
	return &DeviceHandler{Service: service}
}

func (h *DeviceHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ok := checkMethod(w, r, GET)
	if ok {
		ctx := r.Context()
		path := strings.TrimPrefix(r.URL.Path, "/device/")
		deviceID := strings.TrimSuffix(path, "/")
		id, err := strconv.Atoi(deviceID)
		if err != nil {
			returnError(w, "id must be an integer", http.StatusBadRequest)
			return
		}
		device, err := h.Service.GetById(ctx, id)
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
}

func (h *DeviceHandler) UpdateDeviceActivity(w http.ResponseWriter, r *http.Request) {
	ok := checkMethod(w, r, PUT)
	if ok {
		ctx := r.Context()
		var req DeviceActivityRequest
		err := unmarshalRequest(r, &req)
		if err != nil {
			returnError(w, "error in unmarshalling request", http.StatusBadRequest)
			return
		}
		err = h.Service.UpdateDeviceActivity(ctx, req.ID, req.Active)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		returnSuccess(w, nil)
	}
}
