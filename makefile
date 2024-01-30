help:				## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build:			  	## Builds code into a binary
	@go build -o main && zip main.zip main

build-lambda:			  	## Builds code into a binary
	@CGO_ENABLED=0 go build -o bootstrap && zip main.zip bootstrap

run:				## Runs code localy
	@go run .

dev:				## Dev Mode: will restart app on each code changes
	@reflex -r '.go' -s -- go run .


tidy:				## Install && cleans project dependencies in go.mod
	@go mod tidy

test:				## Run Unit tests
	@go test -v ./...

lint:				## Run lint command
	@golangci-lint run  --enable-all -D varnamelen -D interfacer -D maligned -D golint -D scopelint -D exhaustivestruct -D bodyclose -D contextcheck -D nilerr -D noctx -D rowserrcheck -D sqlclosecheck -D structcheck -D tparallel -D wastedassign -D gci -D gosimple -D usestdlibvars -D lll -D paralleltest -D ifshort -D depguard -D varcheck -D deadcode -D nosnakecase -D exhaustruct

lint-fix:				## Run lint command with --fix
	@golangci-lint run  --enable-all -D varnamelen -D interfacer -D maligned -D golint -D scopelint -D exhaustivestruct -D bodyclose -D contextcheck -D nilerr -D noctx -D rowserrcheck -D sqlclosecheck -D structcheck -D tparallel -D wastedassign -D gci -D gosimple -D usestdlibvars -D lll -D paralleltest -D ifshort -D depguard -D varcheck -D deadcode -D nosnakecase -D exhaustruct --fix
