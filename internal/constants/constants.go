package constants

import "fmt"

// LogLevel represents the severity of the log entry.
type LogLevel int

const (
	DEBUG   LogLevel = 1
	INFO    LogLevel = 2
	WARNING LogLevel = 3
	ERROR   LogLevel = 4
)

const (
	DEBUGStr   string = "DEUBG"
	INFOStr    string = "INFO"
	WARNINGStr string = "WARNING"
	ERRORStr   string = "ERROR"
)

var (
	logLevelIntToStr = map[LogLevel]string{
		DEBUG:   DEBUGStr,
		INFO:    INFOStr,
		WARNING: WARNINGStr,
		ERROR:   ERRORStr,
	}

	logLevelStrToInt = map[string]LogLevel{
		DEBUGStr:   DEBUG,
		INFOStr:    INFO,
		WARNINGStr: WARNING,
		ERRORStr:   ERROR,
	}
)

// Converts LogLevel type to string
func LogLevelIntToStr(level LogLevel) string {
	return logLevelIntToStr[level]
}

// Converts String LogLevel to LogLevel type
func LogLevelStrToInt(level string) (LogLevel, error) {
	levelInt, ok := logLevelStrToInt[level]
	if !ok {
		return -1, fmt.Errorf("invalid log level")
	}

	return levelInt, nil
}
