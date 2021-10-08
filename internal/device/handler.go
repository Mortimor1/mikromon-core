package device

import (
	"github.com/Mortimor1/mikromon-core/internal/handlers"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
	logger logging.Logger
}

func NewDeviceHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/devices", h.GetDevices).Methods("GET")
	router.HandleFunc("/devices/{id}", h.GetDeviceById).Methods("GET")
	router.HandleFunc("/devices/{id}", h.CreateDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", h.UpdateDevice).Methods("PUT")
	router.HandleFunc("/devices/{id}", h.DeleteDevice).Methods("DELETE")
}

func (h *handler) GetDevices(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("GetDevices"))
}

func (h *handler) GetDeviceById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("GetDeviceById"))
}

func (h *handler) CreateDevice(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("CreateDevice"))
}

func (h *handler) UpdateDevice(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("UpdateDevice"))
}

func (h *handler) DeleteDevice(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("DeleteDevice"))
}
