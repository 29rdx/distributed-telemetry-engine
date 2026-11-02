package exporter

import (
	"encoding/json"
	"time"
)

type OTLPSpan struct {
	TraceID   string            `json:"trace_id"`
	SpanID    string            `json:"span_id"`
	Name      string            `json:"name"`
	Timestamp int64             `json:"timestamp_unix_nano"`
	Status    string            `json:"status"`
	Tags      map[string]string `json:"attributes"`
}

func FormatSpan(traceID, spanID, name, status string, tags map[string]string) ([]byte, error) {
	span := OTLPSpan{
		TraceID:   traceID,
		SpanID:    spanID,
		Name:      name,
		Timestamp: time.Now().UnixNano(),
		Status:    status,
		Tags:      tags,
	}
	return json.Marshal(span)
}
