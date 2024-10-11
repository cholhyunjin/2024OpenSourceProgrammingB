package main

import (
	"fmt" // c language #include <stdio.h>
)

func main() {
	//var now time.Time = time.Now()
	//month := now.Month()
	//fmt.Println(int(month))
	//fmt.Printf("오늘은 %d년 %d월 %d일 입니다\n", now.Year(), int(now.Month()), now.Day())
	//fmt.Printf("지금 시간은 %d시 %d분 %d초 입니다\n", now.Hour(), now.Minute(), now.Second())
	var army string = "우리는 국가와 국민에 충성을 다하는 대한민국 육군이다."
	armyFixed := string.NewReplacer("!", "국")
	fmt.Println(army)
	fmt.Println(armyFixed.Replace(army))
}
