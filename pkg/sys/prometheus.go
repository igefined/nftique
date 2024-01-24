package sys

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Prometheus struct {
	gatherer prometheus.Gatherer

	requestTotal     *prometheus.CounterVec
	requestsDuration *prometheus.HistogramVec
	requestsInFlight *prometheus.GaugeVec

	defaultUrl string
}

func NewPrometheus(namespace, subsystem, serviceName string) *Prometheus {
	register := prometheus.DefaultRegisterer

	constLabels := make(prometheus.Labels)
	if serviceName != "" {
		constLabels["service"] = serviceName
	}

	counter := promauto.With(register).NewCounterVec(
		prometheus.CounterOpts{
			Name:        prometheus.BuildFQName(namespace, subsystem, "requests_total"),
			Help:        "Total http requests by status code, method and path",
			ConstLabels: constLabels,
		},
		[]string{"status_code", "method", "path"},
	)

	histogram := promauto.With(register).NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        prometheus.BuildFQName(namespace, subsystem, "request_duration_seconds"),
			Help:        "Duration of all HTTP requests by status code, method and path.",
			ConstLabels: constLabels,
			Buckets: []float64{
				0.000000001, 0.000000002, 0.000000005, 0.00000001, 0.00000002, 0.00000005, 0.0000001, 0.0000002, 0.0000005, 0.000001, 0.000002, 0.000005, 0.00001, 0.00002, 0.00005, 0.0001, 0.0002, 0.0005, 0.001, 0.002, 0.005, 0.01, 0.02, 0.05, 0.1, 0.2, 0.5, 1.0, 2.0, 5.0, 10.0, 15.0, 20.0, 30.0, //nolint:lll
			},
		},
		[]string{"status_code", "method", "path"},
	)

	gauge := promauto.With(register).NewGaugeVec(prometheus.GaugeOpts{
		Name:        prometheus.BuildFQName(namespace, subsystem, "requests_in_progress_total"),
		Help:        "All the requests in progress",
		ConstLabels: constLabels,
	}, []string{"method"})

	return &Prometheus{
		gatherer:         prometheus.DefaultGatherer,
		requestTotal:     counter,
		requestsDuration: histogram,
		requestsInFlight: gauge,
		defaultUrl:       "/metrics",
	}
}

func (p *Prometheus) Handler() http.Handler {
	return promhttp.HandlerFor(p.gatherer, promhttp.HandlerOpts{})
}

func (p *Prometheus) FiberMiddleware(ctx *fiber.Ctx) error {
	start := time.Now()
	method := ctx.Route().Method

	if ctx.Route().Path == p.defaultUrl {
		return ctx.Next()
	}

	p.requestsInFlight.WithLabelValues(method).Inc()
	defer func() {
		p.requestsInFlight.WithLabelValues(method).Dec()
	}()

	err := ctx.Next()
	status := fiber.StatusInternalServerError
	if err != nil {
		var e *fiber.Error
		if errors.As(err, &e) {
			status = e.Code
		}
	} else {
		status = ctx.Response().StatusCode()
	}

	path := ctx.Route().Path
	statusCode := strconv.Itoa(status)
	p.requestTotal.WithLabelValues(statusCode, method, path).Inc()

	elapsed := float64(time.Since(start).Nanoseconds()) / 1e9
	p.requestsDuration.WithLabelValues(statusCode, method, path).Observe(elapsed)

	return err
}
