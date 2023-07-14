# Build metadata
ProjectName = $(shell basename "$(PWD)")
GoTestFiles = $(shell go list ./... | grep -v /vendor/)
CompiledFileName=build/$(ProjectName)

## clean: Clean build files. Runs `go clean` internally.
clean:
	@echo "Cleaning build cache"
	@go clean
	@rm -rf build

## compile: Compile the project
compile: clean
	@echo "Compiling project"
	@go build -o $(CompiledFileName) main.go

## unit_test: Run unit tests
unit_test:
	@echo "Running unit tests"
	@go test -v -cover -coverprofile=coverage.out $(GoTestFiles)

## coverage: Run unit tests and generate coverage report
coverage: unit_test
	@echo "Generating coverage report"
	@go tool cover -html=coverage.out

## build_all: Perform cross-compilation for multiple platforms
build_all: clean
	@echo "Performing cross-compilation"
	@gox -osarch="linux/amd64 windows/amd64 darwin/amd64" -output="build/$(ProjectName)-{{.OS}}-{{.Arch}}" .

## help: Show this help message
help: Makefile
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@sed -n 's/^##//p' $<
	@echo
