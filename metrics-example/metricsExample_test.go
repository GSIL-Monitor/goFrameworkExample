package metrics_example

import (
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	c := metrics.NewCounter()
	metrics.Register("foo", c)
	c.Inc(47)

	g := metrics.NewGauge()
	metrics.Register("bar", g)
	g.Update(47)

	r := metrics.NewRegistry()
	g1 := metrics.NewRegisteredFunctionalGauge("cache-evictions", r, func() int64 { return 1 })
	metrics.Register("FunctionalGauge", g1)

	s := metrics.NewExpDecaySample(1028, 0.015) // or metrics.NewUniformSample(1028)
	h := metrics.NewHistogram(s)
	metrics.Register("baz", h)
	h.Update(47)
	h.Update(50)
	h.Update(1000)

	m := metrics.NewMeter()
	metrics.Register("quux", m)
	m.Mark(47)

	t1 := metrics.NewTimer()
	metrics.Register("bang", t1)
	t1.Time(func() {})
	t1.Update(47)

	<-time.After(10 * time.Second)
}
