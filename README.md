# ⚡ Distributed Telemetry Engine
### High-throughput distributed metrics ingestion & real-time telemetry stream engine (Go)

[![CI](https://github.com/29rdx/distributed-telemetry-engine/actions/workflows/ci.yml/badge.svg)](https://github.com/29rdx/distributed-telemetry-engine/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go 1.22](https://img.shields.io/badge/Go-1.22-cyan.svg)](https://go.dev/)

## 📌 Architecture
- **Lock-Free Ring Buffer**: Cache-padded atomic CAS circular queue.
- **Adaptive Tail-Sampling**: Dynamic sampling rate with 100% capture of error traces.
- **Reservoir Quantile Tracking**: Real-time streaming P50, P90, and P99 latency calculation.
- **OTLP Formatter**: OpenTelemetry Protocol (OTLP/JSON) compliant metric serializer.
- **Tag Cardinality Defense**: Bounded tag key-value limiter protecting metric backends.

## 🧪 Testing & Execution
```bash
go test -v -race ./...
```
