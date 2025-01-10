package routing

import (
	"net/http"
	"voting-server/internal/api/routing/routes"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ConnectRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	routes.OpenRoutes(mux)

	return mux
}
