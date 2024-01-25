package sys

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gofiber/fiber/v2"
	cfg "github.com/igefined/nftique/pkg/config"

	"go.uber.org/zap"
)

const (
	HeaderContentType = "Content-Type"
	ApplicationJSON   = "application/json"
)

type Sys struct {
	lg  *zap.Logger
	cfg *cfg.MainCfg

	appInfo    *AppInfo
	prometheus *Prometheus
	httpServer *http.Server
	mux        *http.ServeMux
}

func NewSys(lg *zap.Logger, cfg *cfg.MainCfg, appInfo *AppInfo) *Sys {
	instance := &Sys{
		lg:         lg,
		cfg:        cfg,
		appInfo:    appInfo,
		prometheus: NewPrometheus(cfg.Namespace, cfg.Environment, cfg.ServiceName),
		mux:        http.NewServeMux(),
	}

	instance.initRoutes()

	return instance
}

func (s *Sys) initRoutes() {
	s.mux.Handle("/metrics", s.prometheus.Handler())

	// version
	s.mux.HandleFunc("/version", s.VersionHandler)

	// k8s probes
	s.mux.HandleFunc("/ready", s.ReadyHandler)
	s.mux.HandleFunc("/live", s.LiveHandler)

	// pprof
	s.mux.HandleFunc("/debug/pprof/", pprof.Index)
	s.mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	s.mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	s.mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	s.mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	s.mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
}

func (s *Sys) Start() {
	s.httpServer = &http.Server{
		Addr: fmt.Sprintf(":%s", s.cfg.MonitorPort),
		// TODO from cfg
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 5,
		Handler:      s.mux,
	}

	go s.listenAndServe()
}

func (s *Sys) Stop(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}

	return s.httpServer.Shutdown(ctx)
}

func (s *Sys) PrometheusMiddleware() func(ctx *fiber.Ctx) error {
	return s.prometheus.FiberMiddleware
}

func (s *Sys) listenAndServe() {
	s.lg.Info("sys-http server started serving", zap.String("port", s.cfg.MonitorPort))

	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.lg.Fatal("sys-http server initialize/runtime error", zap.Error(err))
	}
}
