/* string is another builtin type.
   "xxx" ---- string
   'xxx' ---- char
   string is unmutable */

package main

import "fmt"

func main(){
	s := "hello"
	c := []rune(s)          // convert to array
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

//	s1 := "Starting part"
//	+ "Ending part"           // multi lines text: wrong

	s3 := "Starting part" +
	"Ending part"             // multi lines text: right

	s4 := `Starting part
	Ending part`             // original string display

//	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s4)
}
