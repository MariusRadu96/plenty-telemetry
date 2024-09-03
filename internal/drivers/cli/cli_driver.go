package cli

import (
	"encoding/json"
	"fmt"
	"plentytelemetry/internal/constants"
	"plentytelemetry/internal/domain"
	"time"
)

type CLIDriver struct{}

// Create a new CLIDriver instance with the specified filename.
func NewCLIDriver() *CLIDriver {
	return &CLIDriver{}
}

// Log the entry to the command line.
func (d *CLIDriver) Log(entry domain.LogEntry) error {
	attributes, err := json.Marshal(entry.Attributes)
	if err != nil {
		return err
	}

	fmt.Printf("[%s][%s] [%s] %s [Attributes: %s]\n", entry.TraceID, entry.Timestamp.Format(time.RFC3339), constants.LogLevelIntToStr(entry.Level), entry.Message, string(attributes))
	return nil
}
