SHELL = /bin/bash
test-all:
	cd ./service && go test -v ./... -cover
install-client:
	cd client && yarn
clean-up:
	cd ./client && rm -rf ./node_modules
setup-service:
	cd ./service && go mod tidy
start-service:
	cd ./service && go run ./cmd/healthcheck/main.go
start-client:
	cd ./client && yarn start