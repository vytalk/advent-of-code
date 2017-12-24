package main

import (
	"sort"
	"strings"
)

/*
To ensure security, a valid passphrase must contain no duplicate words.

For example:

    aa bb cc dd ee is valid.
    aa bb cc dd aa is not valid - the word aa appears more than once.
    aa bb cc dd aaa is valid - aa and aaa count as different words.

*/
func ValidatePassphrases(passphrases []string) int {
	c := make(chan bool)
	for _, p := range passphrases {
		go validatePassphrase(p, c)
	}

	valid := 0
	for range passphrases {
		if <-c {
			valid++
		}
	}

	return valid
}

func validatePassphrase(passphrase string, c chan bool) {
	strings := strings.Split(passphrase, " ")
	for _, v := range strings {
		count := 0
		for _, _v := range strings {
			if v == _v {
				count++
				if count > 1 {
					c <- false
					return
				}
			}
		}
	}

	c <- true
}

/*
A valid passphrase must contain no two words that are anagrams of each other -
that is, a passphrase is invalid if any word's letters can be rearranged to
form any other word in the passphrase.

For example:

abcde fghij is a valid passphrase.
abcde xyz ecdab is not valid - the letters from the third word can be rearranged
 to form the first word.
a ab abc abd abf abj is a valid passphrase, because all letters need to be used
 when forming another word.
iiii oiii ooii oooi oooo is valid.
oiii ioii iioi iiio is not valid - any of these words can be rearranged to form
 any other word.
Under this new system policy, how many passphrases are valid?
*/
func ValidatePassphrasesB(passphrases []string) int {
	c := make(chan bool)
	for _, p := range passphrases {
		go validatePassphraseB(p, c)
	}

	valid := 0
	for range passphrases {
		if <-c {
			valid++
		}
	}

	return valid
}

func validatePassphraseB(passphrase string, c chan bool) {
	strings := strings.Split(passphrase, " ")
	for _, v := range strings {
		count := 0
		v = sortChars(v)
		for _, _v := range strings {
			_v = sortChars(_v)
			if v == _v {
				count++
				if count > 1 {
					c <- false
					return
				}
			}
		}
	}

	c <- true
}

func sortChars(val string) string {
	split := strings.Split(val, "")
	sort.Strings(split)
	return strings.Join(split, "")
}
