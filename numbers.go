/* int --- 32bit and 64bit depandes on hardware
	-- int32, uint32
	-- type = uint8
   float --- two value: float32 and float64
   They are indepandent, cannot be used fixed */

package main


func main(){
	var a int
	var b int32
	a = 15
	b = a + a      //error  ./numbers.go:18:4: cannot use a + a (type int) as type int32 in assignment

	b = b + 5
}
