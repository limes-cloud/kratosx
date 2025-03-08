package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var _ Observer = (*summary)(nil)

type summary struct {
	sv  *prometheus.SummaryVec
	lvs []string
}

// NewSummary new a prometheus summary and returns Histogram.
func NewSummary(sv *prometheus.SummaryVec) Observer {
	return &summary{
		sv: sv,
	}
}

func (s *summary) With(lvs ...string) Observer {
	return &summary{
		sv:  s.sv,
		lvs: lvs,
	}
}

func (s *summary) Observe(value float64) {
	s.sv.WithLabelValues(s.lvs...).Observe(value)
}
