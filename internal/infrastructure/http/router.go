package http

import (
	"net/http"

	"github.com/elyarsadig/smart-home-iot/internal/infrastructure/http/handler"
)

func NewRouter(roomHandler *handler.RoomHandler, deviceHandler *handler.DeviceHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/device", deviceHandler.GetByName)
	mux.HandleFunc("/room", roomHandler.GetByName)

	return mux
}
