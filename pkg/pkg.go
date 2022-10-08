package pkg

func Dei(s []int, a int) []int {
	l := len(s)
	slice := s[:a]
	for i := a + 1; i < l; i++ {
		slice = append(slice, s[i])
	}
	return slice
}
