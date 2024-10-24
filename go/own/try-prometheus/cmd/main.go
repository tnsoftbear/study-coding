package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reqTotalMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tp_req_total",
			Help: "Request total count",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	inProgressReqMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tp_in_progress_req_count",
		Help: "In-progress request count",
	})

	reqLatencyMetric = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "tp_req_latency",
			Help: "Request latency",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	// Register metrics in Prometheus
	prometheus.MustRegister(reqTotalMetric)
	prometheus.MustRegister(inProgressReqMetric)
	prometheus.MustRegister(reqLatencyMetric)
}

func doHandler(c *gin.Context) {
	startTime := time.Now()

	inProgressReqMetric.Inc()
	defer inProgressReqMetric.Dec()

	// Imitate workload
	delay := time.Duration(rand.Float64()*0.19+0.01) * time.Second
	time.Sleep(delay)

	// Response with random status
	statusCodes := []int{200, 404, 400, 500}
	statusCode := statusCodes[rand.Intn(len(statusCodes))]

	reqTotalMetric.WithLabelValues("GET", "/do", http.StatusText(statusCode)).Inc()

	// Store processing time
	reqLatencyMetric.WithLabelValues("GET", "/do").Observe(time.Since(startTime).Seconds())

	c.String(statusCode, "Done\n")
}

func main() {
	// Metric service
	go func() {
		http.Handle("/", promhttp.Handler())
		http.ListenAndServe(":8000", nil)
	}()

	router := gin.Default()
	router.GET("/do", doHandler)
	router.Run("0.0.0.0:5000")
}
