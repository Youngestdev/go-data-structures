package main

import "fmt"

func maxCharSequence(s string) int {
	if len(s) == 0 { return 0 }
	var (
		maxLen int
		thisLen int = 1
		lastChar rune
	)
	for _, thisChar := range s {
		if thisChar == lastChar {
			thisLen += 1
		} else {
			if maxLen < thisLen {
				maxLen = thisLen
			}
			thisLen = 1
		}
		lastChar = thisChar
	}
	if maxLen < thisLen { maxLen = thisLen }
	return maxLen
}

func main() {
	fmt.Println(maxCharSequence("Azeeez"))
}