package webserver

import (
	"context"
	"github.com/Mortimor1/mikromon-core/internal/config"
	"github.com/Mortimor1/mikromon-core/internal/device"
	"github.com/Mortimor1/mikromon-core/internal/group"
	"github.com/Mortimor1/mikromon-core/internal/subnet"
	"github.com/Mortimor1/mikromon-core/internal/webserver/handlers"
	"github.com/Mortimor1/mikromon-core/internal/webserver/metrics"
	"github.com/Mortimor1/mikromon-core/pkg/client/mongodb"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	logger.Info("Connect to database")
	client, err := mongodb.NewClient(context.TODO(),
		cfg.Database.Mongo.Url,
		"mikromon")
	if err != nil {
		logger.Fatal(err)
	}

	groupRepo := group.NewGroupRepository(client.Collection("group"))
	deviceRepo := device.NewDeviceRepository(client.Collection("device"))
	subnetRepo := subnet.NewSubnetRepository(client.Collection("subnet"))

	// Init http router
	logger.Info("Create new router")
	router := mux.NewRouter()
	router.Use(handlers.Middleware)
	router.Use(handlers.LoggingMiddleware)
	router.Use(handlers.PrometheusMiddleware)

	reg := metrics.NewMetricsRegistry()
	router.Handle("/metrics", promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	groupHandler := group.NewGroupHandler(logger, groupRepo)
	deviceHandler := device.NewDeviceHandler(logger, deviceRepo)
	subnetHandler := subnet.NewSubnetHandler(logger, subnetRepo)

	logger.Info("Register handlers")
	groupHandler.Register(router)
	deviceHandler.Register(router)
	subnetHandler.Register(router)

	s.httpServer = &http.Server{
		Addr:           cfg.Http.BindIp + ":" + cfg.Http.Port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Start http server
	logger.Infof("Server listening on %s:%s", cfg.Http.BindIp, cfg.Http.Port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
