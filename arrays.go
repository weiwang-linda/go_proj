/* Definition: [n]<type>   */

package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])

	a := [3]int{1,2,3}
	b := [...]int{4,5,6,7}
	c := [3][2]int{[2]int{1,2}, [2]int{3,4}, [2]int{5,6}}
	d := [3][2]int{{1,2}, {3,4}, {5,6}}
	fmt.Println(a,b,c,d)
}
