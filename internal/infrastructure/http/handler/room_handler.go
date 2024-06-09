package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/elyarsadig/smart-home-iot/config"
	"github.com/elyarsadig/smart-home-iot/internal/domain/service"
)

type RoomHandler struct {
	Service *service.RoomService
}

func NewRoom(service *service.RoomService) *RoomHandler {
	return &RoomHandler{Service: service}
}

func (h *RoomHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ok := checkMethod(w, r, GET)
	if ok {
		ctx := r.Context()
		path := strings.TrimPrefix(r.URL.Path, "/room/")
		roomID := strings.TrimSuffix(path, "/")
		id, err := strconv.Atoi(roomID)
		if err != nil {
			returnError(w, "id must be an integer", http.StatusBadRequest)
			return
		}
		room, err := h.Service.GetById(ctx, id)
		if err != nil {
			if errors.Is(err, config.NotFoundError) {
				returnError(w, "room not found", http.StatusNotFound)
				return
			}
			returnError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		returnSuccess(w, room)
	}
}
