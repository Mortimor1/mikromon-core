package device

import (
	"context"
	"encoding/json"
	"github.com/Mortimor1/mikromon-core/internal/webserver/handlers"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
	logger *logging.Logger
	r      *deviceRepository
}

func NewDeviceHandler(logger *logging.Logger, r *deviceRepository) handlers.Handler {
	return &handler{
		logger: logger,
		r:      r,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/devices", h.GetDevices).Methods("GET")
	router.HandleFunc("/devices/{id}", h.GetDeviceById).Methods("GET")
	router.HandleFunc("/devices/{ip}", h.GetDeviceByIp).Methods("GET")
	router.HandleFunc("/devices", h.CreateDevice).Methods("POST")
	router.HandleFunc("/devices/{id}", h.UpdateDevice).Methods("PUT")
	router.HandleFunc("/devices/{id}", h.DeleteDevice).Methods("DELETE")
}

func (h *handler) GetDevices(writer http.ResponseWriter, _ *http.Request) {
	d, err := h.r.FindAll(context.Background())
	if err != nil {
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(writer).Encode(d)
}

func (h *handler) GetDeviceById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	d, err := h.r.FindOne(context.Background(), "_id", vars["id"])

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(writer).Encode(d)
}

func (h *handler) GetDeviceByIp(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	d, err := h.r.FindOne(context.Background(), "ipaddress", vars["ip"])

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(writer).Encode(d)
}

func (h *handler) CreateDevice(writer http.ResponseWriter, request *http.Request) {
	var d Device
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&d); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}
	defer request.Body.Close()

	exist, err := h.r.FindOne(context.Background(), "ipaddress", d.IpAddress)

	if exist != nil {
		writer.WriteHeader(http.StatusConflict)
		json.NewEncoder(writer).Encode(map[string]string{"error": "Device already exist"})
		return
	}

	id, err := h.r.Create(context.Background(), &d)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(map[string]string{"id": id})
}

func (h *handler) UpdateDevice(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
	vars := mux.Vars(request)

	var d Device
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&d); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}
	defer request.Body.Close()

	d.Id = vars["id"]

	errUpdate := h.r.Update(context.Background(), &d)
	if errUpdate != nil {
		json.NewEncoder(writer).Encode(map[string]string{"error": errUpdate.Error()})
	}
}

func (h *handler) DeleteDevice(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)

	vars := mux.Vars(request)

	errDelete := h.r.Delete(context.Background(), vars["id"])
	if errDelete != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"error": errDelete.Error()})
	}
}
