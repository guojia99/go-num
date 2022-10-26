package utils

func Contains(s []string, v string) bool {
	for _, n := range s {
		if n == v {
			return true
		}
	}
	return false
}
