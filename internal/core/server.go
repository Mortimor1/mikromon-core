package core

import (
	"context"
	"github.com/Mortimor1/mikromon-core/internal/config"
	"github.com/Mortimor1/mikromon-core/internal/device"
	"github.com/Mortimor1/mikromon-core/internal/group"
	"github.com/Mortimor1/mikromon-core/internal/handlers"
	"github.com/Mortimor1/mikromon-core/pkg/client/mongodb"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	client     *mongo.Database
}

func (s *Server) Run(cfg *config.Config) error {
	// Init logger
	logger := logging.GetLogger()

	// Init DB
	client, err := mongodb.NewClient(context.TODO(),
		"mongodb://mikromon:213WJN8gQ12@10.10.0.10:27017/?authMechanism=DEFAULT",
		"mikromon")
	if err != nil {
		logger.Fatal(err)
	}
	s.client = client

	// Init http router
	logger.Info("create new router")
	router := mux.NewRouter()
	router.Use(handlers.Middleware)
	router.Use(handlers.LoggingMiddleware)

	groupHandler := group.NewGroupHandler(logger)
	deviceHandler := device.NewDeviceHandler(logger)

	logger.Info("register handlers")
	groupHandler.Register(router)
	deviceHandler.Register(router)

	s.httpServer = &http.Server{
		Addr:           cfg.Http.BindIp + ":" + cfg.Http.Port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Start http server
	logger.Infof("server listening on %s:%s", cfg.Http.BindIp, cfg.Http.Port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
