package workspaceservicehttp

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SolomonAIEngineering/backend-core-library/database/redis"
	"github.com/SolomonAIEngineering/backend-core-library/instrumentation"
	"github.com/SolomonAIEngineering/core/api-packages/workspace-service-http/secrets"
	"github.com/gorilla/mux"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	// ErrRouterNotSet is returned when the router is not set for the RestService.
	ErrRouterNotSet = fmt.Errorf("router not set")
	// ErrLoggerNotSet is returned when the logger is not set for the RestService.
	ErrLoggerNotSet = fmt.Errorf("logger not set")
	// ErrInstrumentationNotSet is returned when the instrumentation client is not set for the RestService.
	ErrInstrumentationNotSet = fmt.Errorf("instrumentation client not set")
	// ErrCacheNotSet is returned when the cache client is not set for the RestService.
	ErrCacheNotSet = fmt.Errorf("cache client not set")
	// ErrKMSNotSet is returned when the key management system is not set for the RestService.
	ErrKMSNotSet = fmt.Errorf("key management system not set")
	// ErrDelayConfigNotSet is returned when the delay configuration is not set for the RestService.
	ErrDelayConfigNotSet = fmt.Errorf("delay configuration not set")
	// ErrTracerNotSet is returned when the tracer is not set for the RestService.
	ErrTracerNotSet = fmt.Errorf("tracer not set")
	// ErrHandlerNotSet is returned when the HTTP handler is not set for the RestService.
	ErrHandlerNotSet = fmt.Errorf("handler not set")
)

// Workspace Service HTTP API
// @version 2.0
// @description REST API for managing workspaces and their resources on the Solomon AI platform.
// @title Workspace Service REST API
// @tag workspace-service-rest
// @contact.name Source Code
// @contact.url https://github.com/SolomonAIEngineering/core

// @license.name MIT License
// @license.url https://github.com/SolomonAIEngineering/core/blob/master/LICENSE

// @host workspace-service.platform.svc.cluster.local:9898
// @BasePath /
// @schemes http

// This API provides functionality to create, retrieve, update, and delete workspaces,
// as well as manage associated resources within the Solomon AI ecosystem. It is designed
// to facilitate efficient workspace management for users and administrators.
type RestService struct {
	router          *mux.Router
	logger          *zap.Logger
	handler         http.Handler
	tracer          trace.Tracer
	tracerProvider  *sdktrace.TracerProvider
	instrumentation *instrumentation.Client
	cache           *redis.Client
	kms             secrets.KeyManagement
	delayConfig     *DelayConfig
	serverConfig    *ServerConfig
}

type ServerConfig struct {
	HttpClientTimeout     time.Duration `mapstructure:"http-client-timeout"`
	HttpServerTimeout     time.Duration `mapstructure:"http-server-timeout"`
	ServerShutdownTimeout time.Duration `mapstructure:"server-shutdown-timeout"`
	BackendURL            []string      `mapstructure:"backend-url"`
	UILogo                string        `mapstructure:"ui-logo"`
	UIMessage             string        `mapstructure:"ui-message"`
	UIColor               string        `mapstructure:"ui-color"`
	UIPath                string        `mapstructure:"ui-path"`
	DataPath              string        `mapstructure:"data-path"`
	ConfigPath            string        `mapstructure:"config-path"`
	CertPath              string        `mapstructure:"cert-path"`
	Host                  string        `mapstructure:"host"`
	Port                  string        `mapstructure:"port"`
	SecurePort            string        `mapstructure:"secure-port"`
	PortMetrics           int           `mapstructure:"port-metrics"`
	Hostname              string        `mapstructure:"hostname"`
	H2C                   bool          `mapstructure:"h2c"`
	RandomDelay           bool          `mapstructure:"random-delay"`
	RandomDelayUnit       string        `mapstructure:"random-delay-unit"`
	RandomDelayMin        int           `mapstructure:"random-delay-min"`
	RandomDelayMax        int           `mapstructure:"random-delay-max"`
	RandomError           bool          `mapstructure:"random-error"`
	Unhealthy             bool          `mapstructure:"unhealthy"`
	Unready               bool          `mapstructure:"unready"`
	JWTSecret             string        `mapstructure:"jwt-secret"`
	CacheServer           string        `mapstructure:"cache-server"`
	CacheTlsEnabled       bool          `mapstructure:"cache-tls-enabled"`
}

type DelayConfig struct {
	RandomDelay     bool
	RandomDelayMin  int
	RandomDelayMax  int
	RandomDelayUnit string
	RandomError     bool
}

type IRestService interface {
	// ListenAndServe starts the HTTP server and returns the server instance, the secure server instance, and the number of active connections for each server.
	// The server and secure server are started in the background, and the function returns immediately.
	ListenAndServe() (*http.Server, *http.Server, *int32, *int32)

	// UploadFile is the handler for the /api/v1/file/upload endpoint.
	UploadFile(w http.ResponseWriter, r *http.Request)
}

// RestServiceOption defines a function type used for configuring RestService.
type RestServiceOption func(*RestService)

var _ IRestService = (*RestService)(nil)

// WithRouter configures a *mux.Router for the RestService.
// It is used to set the router that the service will use to register its endpoints.
//
// Parameters:
// - router: A pointer to a mux.Router instance that should be used by the RestService.
//
// Returns:
// - A function that takes a *RestService and sets its router to the provided one.
func WithRouter(router *mux.Router) RestServiceOption {
	return func(s *RestService) {
		s.router = router
	}
}

// WithLogger configures a *zap.Logger for the RestService.
// It allows the service to log messages using the provided logger.
//
// Parameters:
// - logger: A pointer to a zap.Logger instance for logging operations.
//
// Returns:
// - A function that takes a *RestService and sets its logger to the provided one.
func WithLogger(logger *zap.Logger) RestServiceOption {
	return func(s *RestService) {
		s.logger = logger
	}
}

// WithTracer configures a trace.Tracer for the RestService to enable tracing.
// Tracing allows monitoring and observability of requests and operations within the service.
//
// Parameters:
// - tracer: A trace.Tracer instance to use for tracing operations.
//
// Returns:
// - A function that takes a *RestService and sets its tracer to the provided one.
func WithTracer(tracer trace.Tracer) RestServiceOption {
	return func(s *RestService) {
		s.tracer = tracer
	}
}

// WithTracerProvider sets a TracerProvider for the RestService.
// The TracerProvider is responsible for creating Tracers which can then be used to start spans.
//
// Parameters:
// - tracerProvider: A *sdktrace.TracerProvider instance to use for generating tracers.
//
// Returns:
// - A function that takes a *RestService and sets its tracer provider to the provided one.
func WithTracerProvider(tracerProvider *sdktrace.TracerProvider) RestServiceOption {
	return func(s *RestService) {
		s.tracerProvider = tracerProvider
	}
}

// WithInstrumentation configures an instrumentation client for the RestService.
// This client is used to collect and report metrics and traces about the service's operations.
//
// Parameters:
// - instrumentation: An *instrumentation.Client instance for managing service instrumentation.
//
// Returns:
// - A function that takes a *RestService and sets its instrumentation client to the provided one.
func WithInstrumentation(instrumentation *instrumentation.Client) RestServiceOption {
	return func(s *RestService) {
		s.instrumentation = instrumentation
	}
}

// WithCache configures a Redis cache for the RestService.
// This allows the service to use Redis for caching data, improving performance and scalability.
//
// Parameters:
// - cache: A *redis.Client instance for caching operations.
//
// Returns:
// - A function that takes a *RestService and sets its cache to the provided Redis client.
func WithCache(cache *redis.Client) RestServiceOption {
	return func(s *RestService) {
		s.cache = cache
	}
}

// WithKMS configures a Key Management System (KMS) for the RestService.
// This is used for managing cryptographic keys and operations, enhancing the service's security.
//
// Parameters:
// - kms: A secrets.KeyManagement instance for key management operations.
//
// Returns:
// - A function that takes a *RestService and sets its key management system to the provided one.
func WithKMS(kms secrets.KeyManagement) RestServiceOption {
	return func(s *RestService) {
		s.kms = kms
	}
}

// WithDelayConfig configures the delay behavior for the RestService.
// This allows the service to simulate delays and errors for testing and development purposes.
//
// Parameters:
// - config: A *DelayConfig instance containing the delay and error simulation settings.
//
// Returns:
// - A function that takes a *RestService and sets its delay configuration to the provided one.
func WithDelayConfig(config *DelayConfig) RestServiceOption {
	return func(s *RestService) {
		s.delayConfig = config
	}
}

// NewRestService creates a new RestService with the provided options.
// This constructor allows for flexible configuration of the RestService through various options.
//
// Parameters:
// - opts: A variable number of RestServiceOption functions to configure the new RestService instance.
//
// Returns:
// - A pointer to the newly created RestService, configured with the provided options.
func NewRestService(opts ...RestServiceOption) (*RestService, error) {
	s := &RestService{}
	for _, opt := range opts {
		opt(s)
	}

	if s.router == nil {
		s.router = mux.NewRouter()
	}

	// validate the RestService configuration
	if err := s.validateRestService(); err != nil {
		s.logger.Error("failed to validate RestService configuration", zap.Error(err))
		return nil, err
	}

	// register middleware for the RestService
	s.registerMiddlewares()

	// define the handler for the RestService
	if s.serverConfig.H2C {
		s.handler = h2c.NewHandler(s.router, &http2.Server{})
	} else {
		s.handler = s.router
	}

	// print the registered routes for the RestService
	s.printRoutes()

	return s, nil
}

func (s *RestService) validateRestService() error {
	if s.router == nil {
		return ErrRouterNotSet
	}

	if s.logger == nil {
		return ErrLoggerNotSet
	}

	if s.handler == nil {
		return ErrHandlerNotSet
	}

	if s.tracer == nil {
		return ErrTracerNotSet
	}

	if s.instrumentation == nil {
		return ErrInstrumentationNotSet
	}

	if s.cache == nil {
		return ErrCacheNotSet
	}

	if s.kms == nil {
		return ErrKMSNotSet
	}

	if s.delayConfig == nil {
		return ErrDelayConfigNotSet
	}

	return nil
}

func (s *RestService) registerMiddlewares() {
	prom := NewPrometheusMiddleware()
	s.router.Use(prom.Handler)
	otel := NewOpenTelemetryMiddleware()
	s.router.Use(otel)
	httpLogger := NewLoggingMiddleware(s.logger)
	s.router.Use(httpLogger.Handler)
	s.router.Use(versionMiddleware)
	if s.delayConfig.RandomDelay {
		randomDelayer := NewRandomDelayMiddleware(s.delayConfig.RandomDelayMin, s.delayConfig.RandomDelayMax, s.delayConfig.RandomDelayUnit)
		s.router.Use(randomDelayer.Handler)
	}
	if s.delayConfig.RandomError {
		s.router.Use(randomErrorMiddleware)
	}
}

func (s *RestService) printRoutes() {
	s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
}
