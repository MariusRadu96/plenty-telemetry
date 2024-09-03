package file

import (
	"encoding/json"
	"fmt"
	"os"
	"plentytelemetry/internal/constants"
	"plentytelemetry/internal/domain"
	"time"
)

type FileDriver struct {
	file *os.File
}

// Create a new FileDriver instance with the specified filename.
func NewFileDriver(filename string) (*FileDriver, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileDriver{file: file}, nil
}

// Log the entry to the file.
func (d *FileDriver) Log(entry domain.LogEntry) error {
	attributes, err := json.Marshal(entry.Attributes)
	if err != nil {
		return err
	}
	_, err = d.file.WriteString(fmt.Sprintf("[%s][%s] [%s] %s [Attributes: %s]\n", entry.TraceID, entry.Timestamp.Format(time.RFC3339), constants.LogLevelIntToStr(entry.Level), entry.Message, string(attributes)))
	return err
}
