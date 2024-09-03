test-app:
	GOARCH=amd64 go test -p 1 ./... -cover
	rm ./internal/drivers/file/logs_test.log && rm ./logs_test.log