package metrics

import (
	"encoding/json"
	"github.com/Mortimor1/mikromon-core/internal/config"
	"github.com/Mortimor1/mikromon-core/pkg/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	DeviceCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "mikromon_device_count",
		Help: "Device count",
	})

	HttpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

func NewMetricsRegistry() *prometheus.Registry {
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())

	//go metricsRefresh()
	reg.MustRegister(DeviceCount)
	reg.MustRegister(HttpDuration)

	return reg
}

/*
func metricsRefresh() {
	for {
		DeviceCount.Set(float64(getDeviceCount()))
		time.Sleep(1 * time.Minute)
	}
}*/

func GetDeviceCount() int {
	cfg := config.GetConfig()
	logger := logging.GetLogger()
	c := http.Client{Timeout: time.Duration(5) * time.Second}
	resp, err := c.Get("http://localhost:" + cfg.Http.Port + "/devices")
	if err != nil {
		logger.Errorf("Error %s", err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Errorf("Error %s", err)
	}

	var device interface{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		logger.Errorf("Error %s", err)
	}
	resp.Body.Close()
	devices, ok := device.([]interface{})
	if !ok {
		logger.Error("cannot convert the JSON objects")
	}
	return len(devices)
}
