package group

import (
	"encoding/json"
	"github.com/Mortimor1/mikromon-core/internal/handlers"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
	logger *logging.Logger
}

func NewGroupHandler(logger *logging.Logger) handlers.Handler {
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
	json.NewEncoder(writer).Encode(map[string]string{"handler": "not implemented !"})
}

func (h *handler) GetGroupById(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("not implemented !"))
}

func (h *handler) CreateGroup(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("not implemented !"))
}

func (h *handler) UpdateGroup(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("not implemented !"))
}

func (h *handler) DeleteGroup(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("not implemented !"))
}
