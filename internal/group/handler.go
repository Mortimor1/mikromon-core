package group

import (
	"github.com/Mortimor1/mikromon-core/internal/handlers"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
	logger logging.Logger
}

func NewGroupHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/groups", h.GetGroups).Methods("GET")
	router.HandleFunc("/groups/{id}", h.GetGroupById).Methods("GET")
	router.HandleFunc("/groups/{id}", h.CreateGroup).Methods("POST")
	router.HandleFunc("/groups/{id}", h.UpdateGroup).Methods("PUT")
	router.HandleFunc("/groups/{id}", h.DeleteGroup).Methods("DELETE")
}

func (h *handler) GetGroups(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("GetGroups"))
}

func (h *handler) GetGroupById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("GetGroupById"))
}

func (h *handler) CreateGroup(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("CreateGroup"))
}

func (h *handler) UpdateGroup(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("UpdateGroup"))
}

func (h *handler) DeleteGroup(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("DeleteGroup"))
}
