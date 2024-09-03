package http

import (
	"net/http"
	"net/http/httptest"
	"plentytelemetry/internal/constants"
	"plentytelemetry/internal/domain"
	"testing"
	"time"
)

func Test_NewHTTPDriver(t *testing.T) {

	type args struct {
		attributes map[string]string
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{

		{
			name:    "Error attribute",
			wantErr: true,
			args: args{
				attributes: map[string]string{
					"unknown": "test",
				},
			},
		},

		{
			name:    "Success",
			wantErr: false,
			args: args{
				attributes: map[string]string{
					"endpoint": "https://test.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewHTTPDriver(tt.args.attributes)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTTPDriver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_HTTPDriver_Log(t *testing.T) {
	type args struct {
		entry    domain.LogEntry
		mockhttp func() *httptest.Server
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			name: "Error HTTP",
			args: args{
				mockhttp: func() *httptest.Server {
					return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.WriteHeader(http.StatusInternalServerError)
					}))
				},
				entry: domain.LogEntry{
					Timestamp: time.Now(),
					Level:     constants.WARNING,
					Message:   "test",
					Attributes: map[string]interface{}{
						"test": "test",
					},
					TraceID: "test",
				},
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				mockhttp: func() *httptest.Server {
					return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.WriteHeader(http.StatusOK)
					}))
				},
				entry: domain.LogEntry{
					Timestamp: time.Now(),
					Level:     constants.WARNING,
					Message:   "test",
					Attributes: map[string]interface{}{
						"test": "test",
					},
					TraceID: "test",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the mock server
			server := tt.args.mockhttp()
			defer server.Close()

			driver := &httpDriver{endpoint: server.URL}
			err := driver.Log(tt.args.entry)

			if (err != nil) != tt.wantErr {
				t.Errorf("Log() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
