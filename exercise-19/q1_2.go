package main

import "fmt"

func main(){
	i := 0
	loop:   if i < 10 {
			fmt.Printf("%d\n",i)
			i++
			goto loop
		}
}
