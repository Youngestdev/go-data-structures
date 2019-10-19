package recurrences

func Reverse(s string) string {
	// If the length of the string is 1, return the string.
	if len(s) <= 1 { return  s}
	// Recursion sets in.
	return Reverse(s[1:]) + s[:1]
}