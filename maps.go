/* map[<from type>]<to type> */

package main

import "fmt"

func main(){
	monthdays := map[string]int {
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	fmt.Printf("%d\n", monthdays["Dec"])

	year := 0
	for _, days := range monthdays {
		year += days
	}
	fmt.Printf("The whole year have %d days\n", year)

	var value int
	var present bool

	value, present = monthdays["Jan"]
	fmt.Println(value, present)
	v, ok := monthdays["Sep"]
	fmt.Println(v,ok)

	delete(monthdays, "Mar")
	for k, v := range monthdays {
		fmt.Println(k, v)
	}

}
