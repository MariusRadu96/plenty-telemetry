package plentytelemetry

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"plentytelemetry/internal/constants"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func Test_NewLogger(t *testing.T) {
	type args struct {
		filePath string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Error load config",
			args: args{
				filePath: "./no-such-file.yml",
			},
			wantErr: true,
		},

		{
			name: "Success",
			args: args{
				filePath: "./internal/config/config_test.yml",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewLogger(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_log(t *testing.T) {
	timeNow := time.Now()

	monkey.Patch(time.Now, func() time.Time {
		return timeNow
	})
	defer monkey.Unpatch(time.Now)

	type args struct {
		filePath         string
		level            constants.LogLevel
		message, traceID string
		attributes       map[string]interface{}
	}

	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "Log Level To Low",
			args: args{
				filePath: "./internal/config/config_test.yml",
				level:    constants.DEBUG,
				attributes: map[string]interface{}{
					"test": "test",
				},
				message: "test",
				traceID: "test",
			},
			want: "",
		},

		{
			name: "Success",
			args: args{
				filePath: "./internal/config/config_test.yml",
				level:    constants.WARNING,
				attributes: map[string]interface{}{
					"test": "test",
				},
				message: "test",
				traceID: "test",
			},
			want: fmt.Sprintf("[%s][%s] [%s] %s [Attributes: {\"test\":\"test\"}]\n", "test", timeNow.Format(time.RFC3339), constants.WARNINGStr, "test"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.args.filePath)
			assert.NoError(t, err)

			// Save the current Stdout
			oldStdout := os.Stdout

			r, w, _ := os.Pipe()
			os.Stdout = w

			logger.log(tt.args.level, tt.args.message, tt.args.traceID, tt.args.attributes)

			w.Close()
			var buf bytes.Buffer
			io.Copy(&buf, r)

			output := buf.String()

			// Restore the original Stdout
			os.Stdout = oldStdout

			if output != tt.want {
				t.Errorf("log() got = %v, want %v", output, tt.want)
				return
			}

		})
	}
}
