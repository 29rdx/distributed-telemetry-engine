package exporter

import (
	"encoding/json"
	"testing"
)

func TestFormatSpan(t *testing.T) {
	data, err := FormatSpan("trace-1", "span-1", "HTTP GET /api/v1", "OK", map[string]string{"http.status": "200"})
	if err != nil {
		t.Fatalf("failed to format OTLP span: %v", err)
	}
	var res map[string]interface{}
	if err := json.Unmarshal(data, &res); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}
	if res["name"] != "HTTP GET /api/v1" {
		t.Fatalf("unexpected span name: %v", res["name"])
	}
}
