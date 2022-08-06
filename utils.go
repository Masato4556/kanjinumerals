package kanjinumerals

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func merge[T comparable](m ...map[string]T) map[string]T {
	ans := make(map[string]T, 0)
	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}
	}
	return ans
}

func keys[T comparable](m map[string]T) []string {
	ans := make([]string, 0)
	for key, _ := range m {
		ans = append(ans, key)
	}
	return ans
}
