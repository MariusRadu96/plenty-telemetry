package plentytelemetry

import (
	"log"
	"plentytelemetry/internal/config"
	"plentytelemetry/internal/constants"
	"plentytelemetry/internal/domain"
	"plentytelemetry/internal/drivers/cli"
	"plentytelemetry/internal/drivers/file"
	"plentytelemetry/internal/drivers/http"
	"strings"
	"sync"
	"time"
)

// Driver interface that all log drivers must implement.
type Driver interface {
	Log(entry domain.LogEntry) error
}

// Logger is the main struct to manage log entries and drivers.
type logger struct {
	mu       *sync.Mutex
	drivers  []Driver
	logLevel constants.LogLevel
}

// NewLogger creates a new Logger instance.
func NewLogger(filePath string) (*logger, error) {
	config, err := config.LoadConfg(filePath)
	if err != nil {
		return nil, err
	}

	logLevel, err := constants.LogLevelStrToInt(strings.ToUpper(config.LogLevel))
	if err != nil {
		return nil, err
	}

	return &logger{
		mu:       &sync.Mutex{},
		logLevel: logLevel,
		drivers:  initDrivers(config),
	}, nil
}

// Initialize the drivers based on the config file
func initDrivers(conf *config.Config) []Driver {
	drivers := make([]Driver, 0)
	for _, driver := range conf.Drivers {
		switch driver.Type {
		case "cli":
			drivers = append(drivers, cli.NewCLIDriver())
		case "file":
			fileDriver, err := file.NewFileDriver(driver.Attributes)
			if err != nil {
				log.Panic("err initiating file driver:", err)
			}

			drivers = append(drivers, fileDriver)
		case "http":
			httpDriver, err := http.NewHTTPDriver(driver.Attributes)
			if err != nil {
				log.Panic("err initiating http driver:", err)
			}

			drivers = append(drivers, httpDriver)
		}

		// Open to adding new drivers by adding new casses
	}

	return drivers
}

// Log logs a message with the given level and attributes.
func (l *logger) log(level constants.LogLevel, message, traceID string, attributes map[string]interface{}) {
	// Skip logging if log level smaller than the logger's level

	if level < l.logLevel {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	entry := domain.LogEntry{
		Timestamp:  time.Now(),
		Level:      level,
		Message:    message,
		Attributes: attributes,
		TraceID:    traceID,
	}

	for _, driver := range l.drivers {
		driver.Log(entry)
	}

}

// Debug logs a debug message.
func (l *logger) Debug(message, traceID string, attributes map[string]interface{}) {
	l.log(constants.DEBUG, message, traceID, attributes)
}

// Info logs an info message.
func (l *logger) Info(message, traceID string, attributes map[string]interface{}) {
	l.log(constants.INFO, message, traceID, attributes)
}

// Warning logs a warning message.
func (l *logger) Warning(message, traceID string, attributes map[string]interface{}) {
	l.log(constants.WARNING, message, traceID, attributes)
}

// Error logs an error message.
func (l *logger) Error(message, traceID string, attributes map[string]interface{}) {
	l.log(constants.ERROR, message, traceID, attributes)
}
