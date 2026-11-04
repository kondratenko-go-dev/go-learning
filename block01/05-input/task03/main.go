package main

import "fmt"

func main() {
	var word string
	// fmt.Scan stops reading at the first whitespace character.
	fmt.Scan(&word)
	// %q adds double quotes around the string, making hidden spaces or structure visible.
	fmt.Printf("scan result: %q\n", word)
}

/*
What happened to the rest of the string:
When we input "Ivan Petrov", fmt.Scan(&word) only reads "Ivan" because it splits input by whitespaces.
The remaining part of the string (" Petrov\n") stays in the standard input buffer (stdin) unread.

Why %q is more convenient than %s here:
The %q format automatically wraps the output string in double quotes (e.g., "Ivan").
This makes it immediately obvious if there are any trailing spaces, hidden tabs, or newlines
captured by the variable, which %s would just print invisibly.
*/
