package main

func FirstWord(s string) string {
	n := len(s)
	i := 0
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return ""
	}
	start := i
	for i < n && s[i] != ' ' {
		i++
	}
	end := i
	return s[start:end] 
}
