package main

import "fmt"

func LongestWord(s string) string {
	longest := ""
	current := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			current += string(s[i])
		} else {
			if len(current) > len(longest) {
				longest = current
			}
			current = ""
		}
	}

	// check last word (important!)
	if len(current) > len(longest) {
		longest = current
	}

	return longest
}
func main()  {
	fmt.Println(LongestWord("hello world"))     // hello
	fmt.Println(LongestWord("a bb ccc d"))     // ccc
	fmt.Println(LongestWord(""))               // ""
}