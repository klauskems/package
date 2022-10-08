package pkg

// slice 요소 삭제 s는 slice를 받고 a는 지우고 싶은 index
func Dei(s []int, a int) []int {
	l := len(s)
	copy(s[a:], s[a+1:])
	s[l] = 0
	return s
}

// package main

// import (
// 	"fmt"

// 	"github.com/klauskems/package/pkg"
// )

// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6}
// 	slice = pkg.Dei(slice, 3)
// 	fmt.Println(slice)
// }
