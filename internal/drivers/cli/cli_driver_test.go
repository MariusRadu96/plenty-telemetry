package cli

import (
	"plentytelemetry/internal/constants"
	"plentytelemetry/internal/domain"
	"testing"
	"time"
)

func Test_NewCLIDriver_Log(t *testing.T) {
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
			driver := NewCLIDriver()
			err := driver.Log(tt.args.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCLIDriver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
