package group

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
	r      *groupRepository
}

func NewGroupHandler(logger *logging.Logger, r *groupRepository) handlers.Handler {
	return &handler{
		logger: logger,
		r:      r,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/groups", h.GetGroups).Methods("GET")
	router.HandleFunc("/groups/{id}", h.GetGroupById).Methods("GET")
	router.HandleFunc("/groups", h.CreateGroup).Methods("POST")
	router.HandleFunc("/groups/{id}", h.UpdateGroup).Methods("PUT")
	router.HandleFunc("/groups/{id}", h.DeleteGroup).Methods("DELETE")
}

func (h *handler) GetGroups(writer http.ResponseWriter, _ *http.Request) {
	g, err := h.r.FindAll(context.Background())
	if err != nil {
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
	}
	json.NewEncoder(writer).Encode(g)
}

func (h *handler) GetGroupById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	g, err := h.r.FindOne(context.Background(), vars["id"])

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(writer).Encode(g)
}

func (h *handler) CreateGroup(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)

	var g Group
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&g); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}
	defer request.Body.Close()
	id, err := h.r.Create(context.Background(), &g)

	if err != nil {
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
	}

	json.NewEncoder(writer).Encode(map[string]string{"id": id})
}

func (h *handler) UpdateGroup(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
	vars := mux.Vars(request)

	var g Group
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&g); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}
	defer request.Body.Close()

	g.Id = vars["id"]

	errUpdate := h.r.Update(context.Background(), &g)
	if errUpdate != nil {
		json.NewEncoder(writer).Encode(map[string]string{"error": errUpdate.Error()})
	}
}

func (h *handler) DeleteGroup(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)

	vars := mux.Vars(request)

	errDelete := h.r.Delete(context.Background(), vars["id"])
	if errDelete != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"error": errDelete.Error()})
	}
}
