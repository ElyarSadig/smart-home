package http

import (
	"net/http"

	"github.com/elyarsadig/smart-home-iot/internal/infrastructure/http/handler"
)

func NewRouter(roomHandler *handler.RoomHandler, deviceHandler *handler.DeviceHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/device/", deviceHandler.GetById)
	mux.HandleFunc("/room/", roomHandler.GetById)

	return mux
}
