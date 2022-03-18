package main

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	name string
	p    *prometheus.GaugeVec
}

func newVec(name string) *prometheus.GaugeVec {
	labels := []string{"id"}
	p := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: name,
	}, labels)

	prometheus.MustRegister(p)
	return p
}

func isInMetrics(name string) bool {
	for i := 0; i < len(metrics); i++ {
		if metrics[i].name == name {
			return true
		}
	}
	return false
}

func NewMetric(d []Data) {
	for _, data := range d {
		// se registra una nueva metrica en caso no
		// se encuentre en el array global metrics
		// no se registran los duplicados ya que no se admiten duplicados

		if !isInMetrics(data.param) {
			metrics = append(metrics, &Metric{
				name: data.param,
				p:    newVec(data.param),
			})
		}
	}
}
