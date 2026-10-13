package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	rr := httptest.NewRecorder()
	newMux().ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/", nil))
	if rr.Code != http.StatusOK {
		t.Fatalf("oczekiwano 200, jest %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Hello from DevOps Projekt") {
		t.Fatal("brak powitania w odpowiedzi")
	}
}

func TestHealth(t *testing.T) {
	rr := httptest.NewRecorder()
	newMux().ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/health", nil))
	if rr.Code != http.StatusOK || rr.Body.String() != "ok" {
		t.Fatalf("health nie działa: %d %q", rr.Code, rr.Body.String())
	}
}

func TestMetrics(t *testing.T) {
	rr := httptest.NewRecorder()
	newMux().ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/metrics", nil))
	if !strings.Contains(rr.Body.String(), "app_http_requests_total") {
		t.Fatal("brak metryki app_http_requests_total")
	}
}

func TestNotFound(t *testing.T) {
	rr := httptest.NewRecorder()
	newMux().ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/nie-ma", nil))
	if rr.Code != http.StatusNotFound {
		t.Fatalf("oczekiwano 404, jest %d", rr.Code)
	}
}
