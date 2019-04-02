package server

import (
	"net/http"

	"github.com/rrreeeyyy/prometheus-remote-s3-adapter/client"
	"github.com/rrreeeyyy/prometheus-remote-s3-adapter/config"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "A counter for requests to the wrapped handler.",
		},
		[]string{"handler", "code", "method"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "A histogram of latencies for requests.",
			Buckets: []float64{.25, .5, 1, 2.5, 5, 10},
		},
		[]string{"handler", "method"},
	)
	responseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "response_size_bytes",
			Help:    "A histogram of response sizes for requests.",
			Buckets: []float64{200, 500, 900, 1500},
		},
		[]string{"handler"},
	)
)

type Handler struct {
	cfg     *config.Config
	logger  log.Logger
	router  *mux.Router
	readers []client.Reader
}

func instrumentHandler(name string, handlerFunc http.HandlerFunc) http.Handler {
	return promhttp.InstrumentHandlerDuration(
		requestDuration.MustCurryWith(prometheus.Labels{"handler": name}),
		promhttp.InstrumentHandlerCounter(
			requestCounter.MustCurryWith(prometheus.Labels{"handler": name}),
			promhttp.InstrumentHandlerResponseSize(
				responseSize.MustCurryWith(prometheus.Labels{"handler": name}),
				http.HandlerFunc(handlerFunc),
			),
		),
	)
}

func New(logger log.Logger, cfg *config.Config) *Handler {
	router := mux.NewRouter()

	h := &Handler{
		cfg:    cfg,
		logger: logger,
		router: router,
	}

	router.Methods("GET").Path(h.cfg.Web.TelemetryPath).Handler(promhttp.Handler())
	router.Methods("POST").Path("/read").Handler(instrumentHandler("read", h.read))

	return h
}

func (h *Handler) Run() error {
	level.Info(h.logger).Log("ListenAddress", h.cfg.Web.ListenAddress, "msg", "Listening")
	return http.ListenAndServe(h.cfg.Web.ListenAddress, h.router)
}
