package http

import (
	"net/http"

	"github.com/elyarsadig/smart-home-iot/internal/infrastructure/http/handler"
)

func NewRouter(roomHandler *handler.RoomHandler, deviceHandler *handler.DeviceHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/device/", deviceHandler.GetById)
	mux.HandleFunc("/device", deviceHandler.UpdateDeviceActivity)

	mux.HandleFunc("/room/", roomHandler.GetById)
	// mux.HandleFunc("/room/door")
	// mux.HandleFunc("/room/camera")
	// mux.HandleFunc("/room/alarm")

	return mux
}
