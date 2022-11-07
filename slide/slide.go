package slide

func Intersection(a, b []interface{}) bool {
	m := make(map[interface{}]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := m[v]; ok {
			return true
		}
	}
	return false
}

func IntersectionInt(a, b []int) bool {
	m := make(map[int]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := m[v]; ok {
			return true
		}
	}
	return false
}
