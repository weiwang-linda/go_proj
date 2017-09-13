/* control structures */
/* golang has not do and while, it has for, switch, if, select. */

package main

import (
	"fmt"
	"os"
)

func max(a int, b int) {
	/* if statement */
	if a > b {            // '{' is must, same line with if
		fmt.Printf("The max value is: %d\n", a);
	}else {
		fmt.Printf("The max value is: %d\n", b);
	}
}

func writefile(name string) {
	f, err := os.Create(name)
	/* if statement without else statement */
	if err != nil {
		fmt.Println(err)
		return 
	}
	f.WriteString("This is the first file!")
}

func readfile(name string) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := make([]byte, 30)
	n,_ := f.Read(s)
	fmt.Println(string(s[:n]))
}

/* goto statement for loop: 1. tag must be in local function; 2. case sensitive */
func accounter() {
	i := 0
Here:
	if i > 10 {
		goto later
	}
	println(i)
	i++
	goto Here
later:
}

/* for statements:
   1. for init; condition; post { }  //same with for in C language
   2. for condition { }   //same with while
   3. for { }   //dead loop   */

func sum() {
	sum := 0
	for i:=0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sum2() {
	sum, i := 0, 0
	for i < 10 {
		i++
		sum +=i
	}
	fmt.Println(sum)
}

/* golang without ',' statement */
func reverse(a []byte) {
	for i,j:=0, len(a)-1; i<j; i,j = i+1,j-1 {
		a[i], a[j] = a[j], a[i]
	}
}


/*range: is a iteral. Can be used with slice, map, string, array and channel. It will return key-value sets */
func test_range() {
	list := []string{"a","b","c","d","e","f"}      //this is a slice
	for k, v := range list {
		fmt.Printf("The character %s is in the position of %d in the slice.\n",v,k)
	}
	
	for pos, char := range "abc" {     //this is a string
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
}

/*switch struction:
  switch any statement {
  case condition1:
	  return something
  case condition2:
	  return something
  }
  switch {  //equal to if-else-if, no statement means true
  case condition1:
	  return someone
  case condition2:
	  return someone
  }
*/
func test_switch_char(c byte) {
	switch {
	case '0' <= c && c <= '9':
		fmt.Println(c - '0')
	case 'a' <= c && c <= 'f':
		fmt.Println(c - 'a' + 10)
	case 'A' <= c && c <= 'F':
		fmt.Println(c - 'A' + 10)
	}
}

func test_switch_fallthrough(i int) {

	switch i {
	case 0:
	case 1:
		fmt.Println("case 1 is omitted")
	case 2:
		fallthrough
	case 3:
		fmt.Println("case 3 must be executed since using fallthrough.")
	}
}

func test_switch_escape(c byte) bool {
	switch c {
	case ' ','?', '&','=','#','+':
		return true
	}
	return false
}


func test_switch_compare(a,b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] < b[i]:
			return -1
		case a[i] > b[i]:
			return 0
		}
	}

	switch {
		case len(a) < len(b):
			return -1
		case len(a) > len(b):
			return 1
		}
	return 0
}


func main(){
	
	a, b := 20, 25;
	max(a,b)

	/* if statement: set a local variable */
	if path, err := os.Getwd(); err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(path)
	}

	if true && true {
		fmt.Println("true")
	}

	if ! false {
		fmt.Println("true")
	}

	writefile("test.txt")
	readfile("tet.txt")
	readfile("test.txt")

	accounter()
	sum()
	sum2()

//	s := "abcde"      ///TODO: learn how to make a []byte instance
//	reverse(s)

	/* Using break to exit loop */
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	j: for j:=0; j<5; j++ {
		for i:=0;i<10;i++ {
			if i>5 {
				break j
			}
			fmt.Println(i)
		}
	}

	/* using continue to omit other code */
	for i:=0; i < 10; i++ {
		if i > 5 {
			continue
		}
		fmt.Println(i)
	}

	test_range()

	test_switch_char('b')
	test_switch_fallthrough(0)
	test_switch_fallthrough(2)

	ret1 := test_switch_escape('#')
	fmt.Println(ret1)
	ret2 := test_switch_escape('/')
	fmt.Println(ret2)
}
