package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)
	for _, ch := range str1 {
		set1[ch] = true
	}
	for _, ch := range str2 {
		set2[ch] = true
	}
	count := 0
	for ch := range set1 {
		if !set2[ch] {
			count++
		}
	}
	for ch := range set2 {
		if !set1[ch] {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
