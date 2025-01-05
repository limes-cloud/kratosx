package prometheus

import (
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/limes-cloud/kratosx/config"
)

var (
	ins *prom
)

type prom struct {
	counter   map[string]*prometheus.CounterVec
	gauge     map[string]*prometheus.GaugeVec
	histogram map[string]*prometheus.HistogramVec
	summary   map[string]*prometheus.SummaryVec
	mux       sync.RWMutex
}

type Prometheus interface {
	CounterVec(name string) *prometheus.CounterVec
	GaugeVec(name string) *prometheus.GaugeVec
	HistogramVec(name string) *prometheus.HistogramVec
	SummaryVec(name string) *prometheus.SummaryVec
}

func Instance() Prometheus {
	return ins
}

func Init(conf []*config.Prometheus, watcher config.Watcher) {
	if len(conf) == 0 {
		return
	}

	ins = &prom{
		counter:   make(map[string]*prometheus.CounterVec),
		gauge:     make(map[string]*prometheus.GaugeVec),
		histogram: make(map[string]*prometheus.HistogramVec),
		summary:   make(map[string]*prometheus.SummaryVec),
	}
	ins.initFactory(conf)

	watcher("prometheus", func(value config.Value) {
		if err := value.Scan(&conf); err != nil {
			log.Errorf("prometheus配置变更失败：%s", err.Error())
			return
		}
		ins.initFactory(conf)
	})
}

func (p *prom) initFactory(cfs []*config.Prometheus) {
	p.mux.Lock()
	defer p.mux.Unlock()

	for _, cf := range cfs {
		switch cf.Type {
		case "counter":
			p.counter[cf.Name] = prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace: cf.Namespace,
					Subsystem: cf.Subsystem,
					Name:      cf.Name,
					Help:      cf.Help,
				}, cf.Labels)
			prometheus.MustRegister(p.counter[cf.Name])
		case "gauge":
			p.gauge[cf.Name] = prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Namespace: cf.Namespace,
					Subsystem: cf.Subsystem,
					Name:      cf.Name,
					Help:      cf.Help,
				}, cf.Labels)
			prometheus.MustRegister(p.gauge[cf.Name])
		case "histogram":
			p.histogram[cf.Name] = prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Namespace: cf.Namespace,
					Subsystem: cf.Subsystem,
					Name:      cf.Name,
					Help:      cf.Help,
					Buckets:   cf.Buckets,
				}, cf.Labels)
			prometheus.MustRegister(p.histogram[cf.Name])
		case "summary":
			p.summary[cf.Name] = prometheus.NewSummaryVec(
				prometheus.SummaryOpts{
					Namespace:  cf.Namespace,
					Subsystem:  cf.Subsystem,
					Name:       cf.Name,
					Help:       cf.Help,
					Objectives: cf.Objectives,
					MaxAge:     cf.MaxAge,
					AgeBuckets: cf.AgeBuckets,
					BufCap:     cf.BufCap,
				}, cf.Labels)
			prometheus.MustRegister(p.summary[cf.Name])
		}
	}
}

func (p *prom) CounterVec(name string) *prometheus.CounterVec {
	p.mux.RLock()
	defer p.mux.RUnlock()

	return p.counter[name]
}

func (p *prom) GaugeVec(name string) *prometheus.GaugeVec {
	p.mux.RLock()
	defer p.mux.RUnlock()

	return p.gauge[name]
}

func (p *prom) HistogramVec(name string) *prometheus.HistogramVec {
	p.mux.RLock()
	defer p.mux.RUnlock()

	return p.histogram[name]
}

func (p *prom) SummaryVec(name string) *prometheus.SummaryVec {
	p.mux.RLock()
	defer p.mux.RUnlock()

	return p.summary[name]
}
