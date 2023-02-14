package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var jobsDurationHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		//Name:    "jobs_duration_seconds",
		Name: "http_request_duration_seconds_bucket",
		//WithLabelValues: []string{"method", "status"},
		//http_request_duration_seconds_bucket{code="200",handler="/api/v1/token",method="GET",le="+Inf"} 10000
		Help:    "http request duration",
		Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		//Buckets: []float64{1, 2, 5, 10, 20, 60},
	},
	[]string{"job_type"},
)

var (
	httpRequestsResponseTime prometheus.Summary
)

func init() {
	httpRequestsResponseTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "http",
		Name:      "response_time_seconds",
		Help:      "Request response times",
	})

	prometheus.MustRegister(httpRequestsResponseTime)
	prometheus.MustRegister(jobsDurationHistogram)
}

func Middleware(endpoint string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//job.Run()
		//jobsDurationHistogram.WithLabelValues(job.Type()).Observe(duration.Seconds())
		jobsDurationHistogram.WithLabelValues(fmt.Sprintf("%d", http.StatusOK)).Observe(float64(time.Since(start).Seconds()))
		httpRequestsResponseTime.Observe(float64(time.Since(start).Seconds()))

		log.Printf(
			"\033[5m%s \033[0;103m%s\033[0m \033[0;93m%s\033[0m \033[1;44m%s\033[0m",
			r.Method,
			r.RequestURI,
			endpoint,
			time.Since(start),
		)

		next.ServeHTTP(w, r)
	})
}

func main() {

	// 4 buckets, starting from 10 and adding 5 between each
	prometheus.LinearBuckets(10, 5, 4)

	// 4 buckets, starting from 1 and multiplying by 10 between each
	prometheus.ExponentialBuckets(1, 10, 4)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	mux := http.NewServeMux()

	mux.HandleFunc("/api/admin2", func(rw http.ResponseWriter, req *http.Request) {
		delay := int32(r.Float64() * 10000)
		time.Sleep(time.Millisecond * time.Duration(delay))
		fmt.Printf("Responed with 'Hello admin2, prometheus!' in %dms\n", delay)
		rw.Write([]byte("Hello admin 2!"))
	})

	withMetrics := Middleware("/api/admin2", mux)

	mux.HandleFunc("/api/user3", func(rw http.ResponseWriter, req *http.Request) {
		delay := int32(r.Float64() * 5000)
		time.Sleep(time.Millisecond * time.Duration(delay))
		fmt.Printf("pong user 3' in %dms\n", delay)
		rw.Write([]byte("Hello user 3!"))
	})

	withMetrics = Middleware("/api/user", mux)

	mux.Handle("/metrics", promhttp.Handler())

	log.Println("Run api with prometheus in: 5023")
	log.Fatal(http.ListenAndServe(":5023", withMetrics))
}
