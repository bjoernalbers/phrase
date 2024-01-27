// Package lists provides lists of words to generate passphrases from.
package lists

// Lists contains all wordlists grouped by two-letter language code.
var Lists map[string][]string

func init() {
	Lists = make(map[string][]string)

	// TODO: Replace fake data with real wordlist.
	Lists[""] = []string{"correct", "horse", "battery", "staple"}
}
