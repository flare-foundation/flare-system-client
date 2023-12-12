package shared

import (
	"flare-tlc/client/config"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HealthStatus int

const (
	HealthStatusInitializing HealthStatus = 0 // Default prometheus Gauge value, thus it indicates that it was not updated yet
	HealthStatusOk           HealthStatus = 1
	HealthStatusError        HealthStatus = -1
	HealthStatusSyncing      HealthStatus = -2
)

type MetricsBase struct {
	// status onf the client, see HealthStatus constants
	status prometheus.Gauge
}

func NewMetricsBase(namespace string) *MetricsBase {
	m := &MetricsBase{
		status: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "health_status",
			Help:      "Status of the client (0 - initializing, 1 - ok, -1 - error, -2 - syncing)",
		}),
	}
	m.status.Set(float64(HealthStatusInitializing))
	return m
}

func (m *MetricsBase) SetStatus(status HealthStatus) {
	m.status.Set(float64(status))
}

func InitMetricsServer(cfg *config.MetricsConfig) {
	if len(cfg.PrometheusAddress) == 0 {
		return
	}

	r := mux.NewRouter()

	r.Path("/metrics").Handler(promhttp.Handler())
	r.Path("/health").HandlerFunc(healthHandler)

	srv := &http.Server{
		Addr:    cfg.PrometheusAddress,
		Handler: r,
	}
	go func() {
		err := srv.ListenAndServe()
		log.Fatal(err)
	}()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	err := writeHealthResponse(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func writeHealthResponse(w http.ResponseWriter) (err error) {
	ok, err := getHealthStatus()
	if err != nil {
		return
	}
	if ok {
		_, err = w.Write([]byte("true"))
	} else {
		_, err = w.Write([]byte("false"))
	}
	return
}

func getHealthStatus() (bool, error) {
	gatherer := prometheus.DefaultGatherer
	mfs, err := gatherer.Gather()
	if err != nil {
		return false, err
	}

	// Check metrics with suffix "health_status"
	// Status is healthy if all of them have value HealthStatusOk
	for _, mf := range mfs {
		if strings.HasSuffix(mf.GetName(), "health_status") {
			for _, m := range mf.GetMetric() {
				if g := m.GetGauge(); g != nil {
					if g.GetValue() != float64(HealthStatusOk) {
						return false, nil
					}
				}
			}
		}
	}
	return true, nil
}
