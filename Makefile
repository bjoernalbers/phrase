.PHONY: build
build:
	@go build

.PHONY: check
check:
	go run . | grep -Eq "^[a-z]{3,9} [a-z]{3,9} [a-z]{3,9} [a-z]{3,9}$$"
