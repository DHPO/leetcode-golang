package main

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) > 0 && len(p) == 0 {
		return false
	}
	if len(s) == 0 {
        for _, c := range p {
			if c != '*' {
				return false
			}
		}
		return true
    }
	if p[0] == '*' {
		for j := 1; j < len(p) && p[j] != '*'; j ++ {}
		for i := 0; i < len(s); i ++ {
			if isMatch(s[i:], p[1:]) {
				return true
			}
		}
		return false
	} else if p[0] == '?' {
		return isMatch(s[1:], p[1:])
	} else {
		return s[0] == p[0] && isMatch(s[1:], p[1:])
	}
}
