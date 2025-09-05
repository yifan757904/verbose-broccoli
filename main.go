package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total HTTP requests",
	},
	[]string{"method", "path", "status"},
)

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "ping")
	httpRequestsTotal.WithLabelValues(req.Method, req.URL.Path, "200").Inc()
}

func pong(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
	httpRequestsTotal.WithLabelValues(req.Method, req.URL.Path, "400").Inc()
}

func main() {
	reg := prometheus.NewRegistry()
	// reg.MustRegister(prometheus.NewGoCollector())
	reg.MustRegister(pingCounter)
	reg.MustRegister(httpRequestsTotal)

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/pong", pong)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		Registry:          reg,
		EnableOpenMetrics: true,
	}))
	http.ListenAndServe(":8090", nil)
}
