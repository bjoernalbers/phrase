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
	@[[ `$(EXE)` != `$(EXE)` ]]
	# - generate passphrases with 4 lowercase words by default
	@$(EXE) | grep -Eq "^[a-z]{3,9} [a-z]{3,9} [a-z]{3,9} [a-z]{3,9}$$"
	# - exit with error when list does not exist
	@if $(EXE) -l no-such-list 2>/dev/null; then false; fi
