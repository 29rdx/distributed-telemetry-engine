# Telemetry Ingestion Pipeline Architecture

\`\`\`
[Metrics / Traces] ---> [Lock-Free RingBuffer] ---> [Cardinality Limiter]
                                                           |
                      +------------------------------------+
                      |
              [Adaptive Sampler] (100% Error Boost)
                      |
              [Quantile Tracker] (Streaming P50/P90/P99)
                      |
              [OTLP JSON Exporter] ---> OpenTelemetry Collector / Prometheus
\`\`\`
