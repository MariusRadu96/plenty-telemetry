package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"plentytelemetry/internal/domain"
)

type httpDriver struct {
	endpoint string
}

// Create a new FileDriver instance with the specified filename.
func NewHTTPDriver(attributes map[string]string) (*httpDriver, error) {
	endpoint, ok := attributes["endpoint"]

	if !ok {
		return nil, fmt.Errorf("invalid attributes for http driver")
	}

	return &httpDriver{endpoint: endpoint}, nil
}

// Make http entru for the log.
func (d *httpDriver) Log(entry domain.LogEntry) error {
	body, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("failed to marshal telemetry entry: %v", err)
	}

	resp, err := http.Post(d.endpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to send telemetry entry: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received unsuccessful response: %d", resp.StatusCode)
	}

	return err
}
