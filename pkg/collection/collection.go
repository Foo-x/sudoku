package collection

// string

func FilterS(f func(v string) bool, s []string) []string {
	var filtered []string
	for _, x := range s {
		if f(x) {
			filtered = append(filtered, x)
		}
	}
	return filtered
}

func AllS(f func(v string) bool, slice []string) bool {
	for _, x := range slice {
		if !f(x) {
			return false
		}
	}
	return true
}

func ContainsS(slice []string, s string) bool {
	for _, x := range slice {
		if x == s {
			return true
		}
	}
	return false
}

// int

func MaxI(slice []int) int {
	var max int
	for _, x := range slice {
		if max > x {
			max = x
		}
	}
	return max
}
