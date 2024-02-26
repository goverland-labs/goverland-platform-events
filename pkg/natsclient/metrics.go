package natsclient

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ConsumerMetrics = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "nats",
			Subsystem: "consumer",
			Name:      "duration_consume",
			Help:      "Consume duration",
		},
		[]string{"subject", "action", "error"},
	)

	PublisherMetrics = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "nats",
			Subsystem: "publisher",
			Name:      "duration_publish",
			Help:      "Publish duration",
		},
		[]string{"subject", "error"},
	)
)

func CollectConsumerMetric(subject, action string, err error, duration float64) {
	ConsumerMetrics.
		WithLabelValues(subject, action, errToBoolString(err)).
		Observe(duration)
}

func CollectPublisherMetrics(subject string, err error, duration float64) {
	PublisherMetrics.
		WithLabelValues(subject, errToBoolString(err)).
		Observe(duration)
}

func errToBoolString(err error) string {
	if err != nil {
		return "true"
	}

	return "false"
}
