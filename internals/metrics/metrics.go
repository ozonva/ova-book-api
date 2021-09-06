package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type counter struct {
	prometheusCounter prometheus.Counter
}

func New(name, help string) counter {
	return counter{
		prometheusCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: name,
			Help: help,
		}),
	}
}

func (c *counter) Inc() {
	c.prometheusCounter.Inc()
}
