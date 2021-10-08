package core

import (
	"context"
	"github.com/Mortimor1/mikromon-core/internal/device"
	"github.com/Mortimor1/mikromon-core/internal/group"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	logger := logging.GetLogger()

	logger.Info("create new router")
	router := mux.NewRouter()

	groupHandler := group.NewGroupHandler(logger)
	deviceHandler := device.NewDeviceHandler(logger)

	logger.Info("register handlers")
	groupHandler.Register(router)
	deviceHandler.Register(router)

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logger.Infof("server listening on port %s", port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
