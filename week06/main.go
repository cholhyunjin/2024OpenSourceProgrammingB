package main

import (
	"fmt" // c language #include <stdio.h>
	"reflect"
)

func main() {
	i := 13
	var f float32 = 12.9 //f := 12.9
	c1 := "Z"
	c2 := "김"

	fmt.Printf("value i : %d value i : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f) //i * f (mismatched types int and float64)
	//fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))
	fmt.Println(c1, c2)
	fmt.Printf("%x\n", c2)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
}
