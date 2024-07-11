package main

import (
	"log"
	"net/http"
	"url-shortener/cmd/api/application"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Inicializar la aplicación
	application.Initialize()

	// Crear un nuevo registrador de Prometheus
	registry := prometheus.NewRegistry()

	// Crear métricas de ejemplo (puedes crear las tuyas propias)
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)
	requestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests",
		},
		[]string{"method", "path"},
	)

	// Registrar las métricas en el registrador
	registry.MustRegister(httpRequestsTotal)
	registry.MustRegister(requestDuration)

	// Middleware para contar y medir la duración de las solicitudes HTTP
	instrumentedHandler := func(path string, handler http.Handler) http.Handler {
		return promhttp.InstrumentHandlerCounter(httpRequestsTotal.MustCurryWith(prometheus.Labels{"path": path}),
			promhttp.InstrumentHandlerDuration(requestDuration.MustCurryWith(prometheus.Labels{"path": path}), handler))
	}

	// Configurar las rutas
	http.Handle("/", instrumentedHandler("/", http.FileServer(http.Dir("./cmd/api/static"))))
	http.Handle("/shorten", instrumentedHandler("/shorten", http.HandlerFunc(application.URLHandler.ShortenURLHandler)))
	http.Handle("/s/", instrumentedHandler("/s/", http.HandlerFunc(application.URLHandler.RedirectHandler)))

	// Endpoint para exponer las métricas
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	// Iniciar el servidor HTTP
	log.Fatal(http.ListenAndServe(":8080", nil))
}
