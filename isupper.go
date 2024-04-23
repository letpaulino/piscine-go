package piscine

func IsUpper(s string) bool {
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
