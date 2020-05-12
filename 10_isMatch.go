package main

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	if len(p) == 1 {
		return len(s) == 1 && (p == s || p[0] == '.')
	}

	if p[1] != '*' {
		if s == "" {
			return false
		}
		if p[0] == s[0] || p[0] == '.' {
			return isMatch(s[1:], p[1:])
		} else {
			return false
		}
	}

	for {
		if s == "" || (s[0] != p[0] && p[0] != '.') {
			break
		}

		if isMatch(s, p[2:]) {
			return true
		}
		s = s[1:]
	}
	return isMatch(s, p[2:])
}
