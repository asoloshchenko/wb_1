package main

import (
	"fmt"
	"unicode"
)

func main() {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abCdefAaf ", false},
		{"aabcd", false},
		{"", true},
		{"🎈🎆🎇🧨✨", true},
		{"🎈🎆🎈", false},
		{"123321", false},
	}

	for _, tc := range testCases {
		result := IsAllRuneUnique(tc.input)
		if result != tc.expected {
			fmt.Printf("IsAllRuneUnique(%q) = %v; want %v\n", tc.input, result, tc.expected)
		}
	}
}

func IsAllRuneUnique(s string) bool {
	temp := make(map[rune]struct{})
	var lowerRune rune
	for _, v := range s {
		lowerRune = unicode.ToLower(v)
		if _, ok := temp[lowerRune]; ok {
			return false
		}
		temp[lowerRune] = struct{}{}
	}
	return true
}
