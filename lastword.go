package main

func LastWord(s string) string {
	n := len(s)
	i := n - 1

	for i >= 0 && s[i] == ' ' {
		i--
	}

	if i < 0 {
		return ""
	}

	end := i + 1

	for i >= 0 && s[i] != ' ' {
		i--
	}

	start := i + 1

	return s[start:end]
}