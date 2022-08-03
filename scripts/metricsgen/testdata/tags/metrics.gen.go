// Code generated by metricsgen. DO NOT EDIT.

package tags

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		WithLabels: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "with_labels",
			Help:      "",
		}, append(labels, "step", "time")).With(labelsAndValues...),
		WithExpBuckets: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "with_exp_buckets",
			Help:      "",

			Buckets: stdprometheus.ExponentialBuckets(.1, 100, 8),
		}, labels).With(labelsAndValues...),
		WithBuckets: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "with_buckets",
			Help:      "",

			Buckets: []float64{1, 2, 3, 4, 5},
		}, labels).With(labelsAndValues...),
		Named: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "metric_with_name",
			Help:      "",
		}, labels).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		WithLabels:     discard.NewCounter(),
		WithExpBuckets: discard.NewHistogram(),
		WithBuckets:    discard.NewHistogram(),
		Named:          discard.NewCounter(),
	}
}
