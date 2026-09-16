package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Привет, Go! 🚀"

	bytesLen := len(s)
	fmt.Printf("bytes: %d\n", bytesLen)

	runesLen := utf8.RuneCountInString(s)
	fmt.Printf("runes: %d\n", runesLen)

	fmt.Printf("first byte: %d\n", s[0])

	runes := []rune(s)
	fmt.Printf("first rune: %c\n", runes[0])

	fmt.Printf("last rune: %c\n", runes[len(runes)-1])

	sWithoutRocket := "Привет, Go! "
	rocketSize := len(s) - len(sWithoutRocket)
	fmt.Printf("rocket size in bytes: %d\n", rocketSize)
}

/*
Why len(s) is not suitable for counting characters:
In Go, a string is a read-only slice of bytes, and the built-in len(s) function returns
the number of bytes, not characters (runes). Since Go encodes strings using UTF-8,
characters can take anywhere from 1 to 4 bytes. For instance, Cyrillic letters take 2 bytes
and emojis take 4 bytes. Therefore, len(s) will overestimate the actual character count
for any string containing non-ASCII characters.

The only case where len(s) would match the character count:
len(s) accurately represents the number of characters if and only if the string consists
entirely of ASCII characters (standard English letters, digits, and basic punctuation).
In the ASCII standard, every character is encoded using exactly 1 byte.
*/
