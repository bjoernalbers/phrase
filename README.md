# phrase - the passphrase generator

`phrase` is a command-line tool to generate easy-to-remember passwords from
random words.

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
    correct horse battery staple

Getting help:

    $ phrase -h
    ...

## License

phrase is released under the [MIT License](LICENSE).

The built-in wordlists come from these sources:

- [`de`](https://github.com/bjoernalbers/diceware-wordlist-german/blob/main/wordlist-german-diceware.txt): Copyright by [Björn Albers](https://github.com/bjoernalbers/diceware-wordlist-german) ([MIT License](https://github.com/bjoernalbers/diceware-wordlist-german/blob/main/LICENSE))
- [`en`](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt): Copyright by [Electronic Frontier Foundation](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) ([Creative Commons Attribution License](https://www.eff.org/copyright))
- [`nl`](https://mko.re/diceware/diceware-wordlist-composites-nl.txt): Copyright by [Remko Tronçon](https://el-tramo.be/blog/diceware-nl) ([MIT License](https://github.com/remko/dicewords/blob/master/LICENSE))
