# phrase - the passphrase generator

`phrase` is a command-line tool to generate passphrases, that is
passwords made from random words.

## Installation

### Download

Just download the
[latest release](https://github.com/bjoernalbers/phrase/releases/latest)
for your platform and make it executable, i.e. like this:

    $ curl -L https://github.com/bjoernalbers/phrase/releases/latest/download/phrase-darwin-arm64 -o /usr/local/bin/phrase
    $ chmod +x /usr/local/bin/phrase

### Build it yourself

Clone this repo and build the binary via `make`.

## Usage

Generate random passphrase:

    $ phrase
    sattel metapher dorn mechanik

Getting help:

    $ phrase -h
    ...

## Origin of wordlists

The built-in wordlists (currently only german) come from these sources:

- [de](passphrase/de.txt): [bjoernalbers/diceware-wordlist-german](https://github.com/bjoernalbers/diceware-wordlist-german)
