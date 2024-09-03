package domain

import (
	"plentytelemetry/internal/constants"
	"time"
)

// LogEntry represents a single log entry with metadata.
type LogEntry struct {
	Timestamp  time.Time              `json:"timestamp"`
	Level      constants.LogLevel     `json:"level"`
	Message    string                 `json:"message"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	TraceID    string                 `json:"trace_id,omitempty"`
}
