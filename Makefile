.PHONY: build
build:
	@go build

.PHONY: check
check:
	[[ $$(go run .) !=  $$(go run .) ]]
	go run . | grep -Eq "^[a-z]{3,9} [a-z]{3,9} [a-z]{3,9} [a-z]{3,9}$$"
