package comma

import "strings"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		for i := len(s1) - 1; i >= 0; i-- {
			if !strings.Contains(s2, s1[i:i]) {
				return false
			} else {
				s1 = s1[:strings.LastIndex(s2, s1[i:i])] +
					s1[strings.LastIndex(s2, s1[i:i])+1:]
			}
		}
	}
	return true
}