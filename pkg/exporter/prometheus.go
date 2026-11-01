package exporter

import (
	"fmt"
	"github.com/29rdx/distributed-telemetry-engine/pkg/collector"
)

func FormatPrometheus(m collector.MetricSnapshot) string {
	return fmt.Sprintf(
		"# HELP telemetry_alloc_bytes Current allocated memory\n"+
			"# TYPE telemetry_alloc_bytes gauge\n"+
			"telemetry_alloc_mb{node=\"agent-01\"} %.2f\n"+
			"# HELP telemetry_goroutines Active goroutines count\n"+
			"# TYPE telemetry_goroutines gauge\n"+
			"telemetry_goroutines{node=\"agent-01\"} %d\n",
		m.AllocMB, m.Goroutines,
	)
}
