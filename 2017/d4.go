package main

import (
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
			}
		}

		if count > 1 {
			c <- false
			return
		}
	}

	c <- true
}
