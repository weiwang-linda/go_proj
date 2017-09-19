/* slice type is reference */

package main

import "fmt"

func main(){
	fmt.Printf("Make a slice of '0'\n")
	sl := make([]int, 10)
	fmt.Println(sl)

	fmt.Printf("new a slice from an array or a slice\n")
	a := [...]int{1,2,3,4,5}
	s1 := a[2:4]
	s2 := a[1:5]
	s3 := a[:]
	s4 := a[:4]
	s5 := s2[:]
	fmt.Println(s1,s2,s3,s4,s5)

	fmt.Printf("Error case: overflow\n")
	var array[100]int
	slice := array[0:99]
	slice[98] = 'a'
//	slice[99] = 'a'

	fmt.Printf("Append slice\n")
	ss0 := []int{0,0}
	ss1 := append(ss0, 2)
	ss2 := append(ss1, 3,5,7)
	ss3 := append(ss2, ss0...)
	fmt.Println(ss0,ss1,ss2,ss3)

	fmt.Printf("Copy\n")
	var aa = [...]int{0,1,2,3,4,5,6,7}
	var src = make([]int, 6)
	n1 := copy(src, aa[0:])
	fmt.Println(n1, src)
	n2 := copy(src, src[2:])
	fmt.Println(n2,src)
}
