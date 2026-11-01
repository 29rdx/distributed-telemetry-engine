package collector

import (
	"runtime"
	"time"
)

type MetricSnapshot struct {
	Timestamp   int64   `json:"timestamp"`
	AllocMB     float64 `json:"alloc_mb"`
	TotalAlloc  float64 `json:"total_alloc_mb"`
	NumGC       uint32  `json:"num_gc"`
	Goroutines  int     `json:"goroutines"`
}

func SampleSystemMetrics() MetricSnapshot {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return MetricSnapshot{
		Timestamp:   time.Now().UnixMilli(),
		AllocMB:     float64(m.Alloc) / 1024 / 1024,
		TotalAlloc:  float64(m.TotalAlloc) / 1024 / 1024,
		NumGC:       m.NumGC,
		Goroutines:  runtime.NumGoroutine(),
	}
}
