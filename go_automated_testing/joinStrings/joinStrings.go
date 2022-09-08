package joinstrings

import "strings"

/*
Function JoinWithCommas , we use the slice operator to gather every
phrase in the slice except the last, and pass them to
strings.Join to join them together in a single string, with a
comma and a space between each. Then we add the word
and (surrounded by spaces), and end the string with the
ﬁnal phrase.
*/

func JoinWithCommas(phrases []string) string {
	// Join every phrase except the last with commas
	results := strings.Join(phrases[:len(phrases)-1], ", ")

	// Insert a comma and the word and before the last phrase
	results += ", and "

	// Add the last phrase
	results += phrases[len(phrases)-1]

	// return the joined string
	return results
}
