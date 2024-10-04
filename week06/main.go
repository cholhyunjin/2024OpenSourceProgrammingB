package main

import (
	"fmt" // c language #include <stdio.h>
)

func main() {
	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, b, s, i)
	fmt.Printf("%f%t%s%d\n", f, b, s, i)

	i = 3
	f = 2.7
	fmt.Print("\n\n", f <= float64(i), "\n")

	//i := 13
	//var f float32 = 12.9 //f := 12.9
	//c1 := "Z"
	//c2 := "김"

	//fmt.Printf("value i : %d value i : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f) //i * f (mismatched types int and float64)
	//fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	//fmt.Printf("%d * %f = %d\n", i, f, i*int(f))
	//fmt.Println(c1, c2)
	//fmt.Printf("%x\n", c2) // 유니코드 '김'에 대한 16진수 출력
	//fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i),reflect.TypeOf(c1),reflect.TypeOf(c2))
}
