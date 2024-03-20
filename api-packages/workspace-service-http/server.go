package workspaceservicehttp

import (
	"fmt"
	"net/http"
	"path"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func (s *RestService) startServer() *http.Server {
	// determine if the port is specified
	if s.serverConfig.Port == "0" {
		// move on immediately
		return nil
	}

	srv := &http.Server{
		Addr:         s.serverConfig.Host + ":" + s.serverConfig.Port,
		WriteTimeout: s.serverConfig.HttpServerTimeout,
		ReadTimeout:  s.serverConfig.HttpServerTimeout,
		IdleTimeout:  2 * s.serverConfig.HttpServerTimeout,
		Handler:      s.handler,
	}

	// start the server in the background
	go func() {
		s.logger.Info("Starting HTTP Server.", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatal("HTTP server crashed", zap.Error(err))
		}
	}()

	// return the server and routine
	return srv
}

func (s *RestService) startSecureServer() *http.Server {
	// determine if the port is specified
	if s.serverConfig.SecurePort == "0" {
		// move on immediately
		return nil
	}

	srv := &http.Server{
		Addr:         s.serverConfig.Host + ":" + s.serverConfig.SecurePort,
		WriteTimeout: s.serverConfig.HttpServerTimeout,
		ReadTimeout:  s.serverConfig.HttpServerTimeout,
		IdleTimeout:  2 * s.serverConfig.HttpServerTimeout,
		Handler:      s.handler,
	}

	cert := path.Join(s.serverConfig.CertPath, "tls.crt")
	key := path.Join(s.serverConfig.CertPath, "tls.key")

	// start the server in the background
	go func() {
		s.logger.Info("Starting HTTPS Server.", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServeTLS(cert, key); err != http.ErrServerClosed {
			s.logger.Fatal("HTTPS server crashed", zap.Error(err))
		}
	}()

	// return the server
	return srv
}

func (s *RestService) startMetricsServer() {
	if s.serverConfig.PortMetrics > 0 {
		mux := http.DefaultServeMux
		mux.Handle("/metrics", promhttp.Handler())
		mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%v", s.serverConfig.PortMetrics),
			Handler: mux,
		}

		srv.ListenAndServe()
	}
}
