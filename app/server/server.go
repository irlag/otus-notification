package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"otus-notification/app/metrics"
	"otus-notification/app/services"

	eventApi "otus-notification/app/api/event"
	appProcessors "otus-notification/app/processors"

	"otus-notification/app/api"

	"otus-notification/app/config"
)

type Server struct {
	config      *config.Config
	Prometheus  *prometheus.Registry
	Logger      *zap.Logger
	Router      *mux.Router
	Services    *services.Services
	HttpMetrics *HttpMetrics
}

func New(
	config *config.Config,
	log *zap.Logger,
	processors *appProcessors.Processors,
) *Server {
	server := &Server{
		config: config,
		Logger: log,
	}

	server.configurePrometheus()
	server.initializeMetrics()

	appMetrics := metrics.New()
	appMetrics.MustRegisterMetrics(server.Prometheus)

	server.Router = NewRouter()

	server.Services = services.New(log, config)

	api.NewMetricsApi(server.Prometheus).HandleMethods(server.Router)
	api.NewHealthcheckApi(processors).HandleMethods(server.Router)
	eventApi.NewEventApi(processors).HandleMethods(server.Router)

	return server
}

func (s *Server) Start() error {
	url := fmt.Sprintf("%s:%s", s.config.BindAddress, s.config.Port)

	s.Logger.Info(fmt.Sprintf("starting api server at %s", url))

	corsAllowOrigin := handlers.AllowedOrigins([]string{"*"})

	middlewares := NewMiddlewares(s.Services, s.HttpMetrics)

	s.Router.Use(
		middlewares.StartedAtMiddleware(),
		middlewares.ResponseMiddleware(),
	)

	return http.ListenAndServe(url,
		handlers.CORS(corsAllowOrigin)(
			middlewares.ContentTypeApplicationJsonMiddleware(
				handlers.CompressHandler(
					handlers.LoggingHandler(os.Stdout, s.Router),
				),
			),
		),
	)
}
