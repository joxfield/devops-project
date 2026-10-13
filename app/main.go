// Prosta aplikacja webowa w Go z endpointem /metrics dla Prometheusa.
// Bez zewnętrznych zależności – metryki w formacie tekstowym Prometheusa.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var (
	version      = getenv("APP_VERSION", "dev")
	requestsOK   atomic.Int64
	requestsErr  atomic.Int64
	startTime    = time.Now()
)

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		requestsErr.Add(1)
		http.NotFound(w, r)
		return
	}
	requestsOK.Add(1)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!doctype html><html><head><title>DevOps Projekt</title></head>
<body style="font-family:sans-serif;text-align:center;margin-top:15vh">
<h1>Hello from DevOps Projekt!</h1>
<p>Wersja: <b>%s</b></p>
<p>Host: %s</p>
</body></html>`, version, hostname())
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func metricsHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	fmt.Fprintf(w, "# HELP app_http_requests_total Liczba zapytań HTTP.\n")
	fmt.Fprintf(w, "# TYPE app_http_requests_total counter\n")
	fmt.Fprintf(w, "app_http_requests_total{status=\"ok\"} %d\n", requestsOK.Load())
	fmt.Fprintf(w, "app_http_requests_total{status=\"error\"} %d\n", requestsErr.Load())
	fmt.Fprintf(w, "# HELP app_uptime_seconds Czas działania aplikacji.\n")
	fmt.Fprintf(w, "# TYPE app_uptime_seconds gauge\n")
	fmt.Fprintf(w, "app_uptime_seconds %.0f\n", time.Since(startTime).Seconds())
	fmt.Fprintf(w, "# HELP app_info Informacje o wersji.\n")
	fmt.Fprintf(w, "# TYPE app_info gauge\n")
	fmt.Fprintf(w, "app_info{version=%q} 1\n", version)
}

func hostname() string {
	h, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return h
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/metrics", metricsHandler)
	return mux
}

func main() {
	port := getenv("PORT", "8080")
	log.Printf("Start aplikacji (wersja %s) na porcie %s", version, port)
	log.Fatal(http.ListenAndServe(":"+port, newMux()))
}
