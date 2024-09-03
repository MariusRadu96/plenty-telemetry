package file

import (
	"plentytelemetry/internal/constants"
	"plentytelemetry/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewFileDriver(t *testing.T) {

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{

		{
			name:    "Error file",
			wantErr: true,
			args: args{
				path: "./invalid_path/logs_test.log",
			},
		},

		{
			name:    "Success",
			wantErr: false,
			args: args{
				path: "./logs_test.log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewFileDriver(tt.args.path)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewCLIDriver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_NewFileDriver_Log(t *testing.T) {
	timeNow := time.Now()

	type args struct {
		entry domain.LogEntry
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{

		{
			name:    "Success",
			wantErr: false,
			args: args{
				entry: domain.LogEntry{
					Timestamp: timeNow,
					Level:     constants.WARNING,
					Message:   "test",
					Attributes: map[string]interface{}{
						"test": "test",
					},
					TraceID: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver, err := NewFileDriver("logs_test.log")
			assert.NoError(t, err)
			err = driver.Log(tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFileDriver_Log() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
