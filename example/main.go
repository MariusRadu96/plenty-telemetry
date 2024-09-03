package main

import (
	"fmt"
	"log"
	"plentytelemetry"

	"github.com/google/uuid"
)

func main() {

	logger, err := plentytelemetry.NewLogger("./config.yml")
	if err != nil {
		log.Panicf("err initiating logger: %s", err)
	}

	logger.Debug("This is a debug message, not logged due to log level being INFO", uuid.NewString(), nil)
	logger.Info("Starting application...", uuid.NewString(), map[string]interface{}{"Environment": "development"})
	logger.Warning("Low disk space warning!", uuid.NewString(), map[string]interface{}{"DiskSpace": "500MB"})
	logger.Error("Failed to connect to database", uuid.NewString(), map[string]interface{}{"Database": "users_db", "ErrorCode": 1001})

	// Logging with attributes and TraceID
	logger.Info("User profile updated", uuid.NewString(), map[string]interface{}{"UserID": 123, "Operation": "Update"})
	logger.Error("Payment transaction failed", uuid.NewString(), map[string]interface{}{"UserID": 456, "TransactionID": "789", "Amount": 100})

	fmt.Println("Logs have been recorded.")
}
