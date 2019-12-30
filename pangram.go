package pangram

import (
	"strings"
	"unicode"
)

// IsPangram is a sentence using every letter of the alphabet at least once.
func IsPangram(sentence string) (ans bool) {

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charList := strings.Split(letters, "")
	charMap := make(map[string]bool, 26)

	// make a map of all ascii chars
	for _, ch := range charList {
		charMap[ch] = false
	}
	// fmt.Println("charMap", charMap)

	// set map to true if ch in sentence
	for _, ch := range sentence {
		ch = unicode.ToUpper(ch)
		charMap[string(ch)] = true
	}
	// fmt.Println("charMap", charMap)

	// if a single false exists (not in sentence, bail)
	for _, ok := range charMap {
		if ok == false {
			return false
		}
	}

	return true
}
