package greet

func Hello(name string) string {
	return whisper("Hello, " + name)
}

func whisper(text string) string {
	return text + "..."
}
