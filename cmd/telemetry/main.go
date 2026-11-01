package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/29rdx/distributed-telemetry-engine/pkg/collector"
	"github.com/29rdx/distributed-telemetry-engine/pkg/exporter"
)

func main() {
	fmt.Println("🚀 Distributed Telemetry Engine v1.0.0 initializing...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			metrics := collector.SampleSystemMetrics()
			_ = exporter.FormatPrometheus(metrics)
		}
	}()

	fmt.Println("⚡ Telemetry agent running. Press Ctrl+C to terminate.")
	<-stop
	fmt.Println("🛑 Graceful shutdown complete.")
}
