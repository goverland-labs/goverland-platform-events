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
			Buckets:   []float64{.005, .01, .025, .05, .075, .1, .15, .2, .25, .5, 1, 2.5, 5, 10, 15, 30},
		},
		[]string{"subject", "action", "error"},
	)

	PublisherMetrics = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "nats",
			Subsystem: "publisher",
			Name:      "duration_publish",
			Help:      "Publish duration",
			Buckets:   []float64{.005, .01, .025, .05, .075, .1, .15, .2, .25, .5, 1, 2.5, 5, 10, 15, 30},
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
