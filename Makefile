EXE := ./phrase
SRC := $(shell find . -name '*.go' -or -name 'go.mod')

$(EXE): $(SRC)
	@go build

.PHONY: check
check: $(EXE)
	# phrase should...
	# - display usage instructions
	@$(EXE) -h 2>&1 | grep -q ^Usage
	# - return each time a different passphrase
	@[[ $$(go run .) !=  $$(go run .) ]]
	# - generate passphrases with 4 lowercase words by default
	@go run . | grep -Eq "^[a-z]{3,9} [a-z]{3,9} [a-z]{3,9} [a-z]{3,9}$$"
