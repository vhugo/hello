package hello

func Hello(s string) string {
	if s != "" {
		return "Hello, " + s + "!"
	}
	return "Hello, World!"
}
