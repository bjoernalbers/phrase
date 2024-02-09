EXE := ./phrase
SRC := $(shell find . -name '*.go' -or -name 'go.mod')

$(EXE): $(SRC)
	@go build

.PHONY: unit
unit:
	@go test ./...

.PHONY: integration
integration: $(EXE)
	@go test integration_test.go
