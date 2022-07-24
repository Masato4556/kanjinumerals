package kanjinumerals

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
