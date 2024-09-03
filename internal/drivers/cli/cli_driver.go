package cli

import (
	"encoding/json"
	"fmt"
	"plentytelemetry/internal/constants"
	"plentytelemetry/internal/domain"
	"time"
)

type cliDriver struct{}

// Create a new cliDriver instance with the specified filename.
func NewCLIDriver() *cliDriver {
	return &cliDriver{}
}

// Log the entry to the command line.
func (d *cliDriver) Log(entry domain.LogEntry) error {
	attributes, err := json.Marshal(entry.Attributes)
	if err != nil {
		return err
	}

	fmt.Printf("[%s][%s] [%s] %s [Attributes: %s]\n", entry.TraceID, entry.Timestamp.Format(time.RFC3339), constants.LogLevelIntToStr(entry.Level), entry.Message, string(attributes))
	return nil
}
