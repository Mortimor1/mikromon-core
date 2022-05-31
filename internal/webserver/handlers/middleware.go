package handlers

import (
	"github.com/Mortimor1/mikromon-core/internal/config"
	"github.com/Mortimor1/mikromon-core/internal/webserver/metrics"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	logger := logging.GetLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.GetConfig()
		if cfg.Debug {
			logger.Debug("method: ", r.Method,
				", url: ", r.RequestURI)
		}
		next.ServeHTTP(w, r)
	})
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(metrics.HttpDuration.WithLabelValues(path))
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)
		timer.ObserveDuration()

		if strings.Contains(path, "metrics") {
			metrics.DeviceCount.Set(float64(metrics.GetDeviceCount()))
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
