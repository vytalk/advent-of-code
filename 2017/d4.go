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
	valid := 0
	for _, p := range passphrases {
		local := 0
		for _, v := range strings.Split(p, " ") {
			if strings.Count(p, v) > 1 {
				local++
			}
		}
		if local%2 == 0 {
			valid++
		}
	}

	return valid
}
