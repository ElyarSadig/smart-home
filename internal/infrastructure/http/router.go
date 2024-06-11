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
	mux.HandleFunc("/room", roomHandler.UpdateRoomStatus)
	mux.HandleFunc("/room/energy-saving/", roomHandler.GetRoomEnergySaving)

	return mux
}
