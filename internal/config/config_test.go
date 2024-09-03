package config

import (
	"os"
	"testing"
)

// Helper function
func createTempFile(t *testing.T, content string) *os.File {
	file, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	if _, err := file.Write([]byte(content)); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}
	return file
}

func Test_LoadConfig(t *testing.T) {
	validYAML := `
log_level: "warn"
drivers:
  - type: "file"
    file_path: "./logs_test.log"
  - type: "cli"
`

	invalidYAML := `
log_level: "warn"
drivers:
  - type: "file"
    file_path: "./logs_test.log"
	extra_field: "invalid"    
`

	type args struct {
		fileContent string
	}

	tests := []struct {
		name    string
		want    *Config
		wantErr bool
		args    args
	}{
		{
			name:    "Invalid Config File",
			wantErr: true,
			args: args{
				fileContent: invalidYAML,
			},
		},

		{
			name:    "Success",
			wantErr: false,
			args: args{
				fileContent: validYAML,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := createTempFile(t, tt.args.fileContent)
			defer os.Remove(file.Name())
			filePath := file.Name()

			_, err := LoadConfg(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
