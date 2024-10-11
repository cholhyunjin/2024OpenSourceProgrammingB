package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("점수 입력 : ")
	i := bufio.NewReader(os.Stdin)
	score, err := i.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	score = (strings.TrimSpace(score))              //줄바꿈, 띄어쓰기, 탭 등 제거 (python strip과 유사)
	realscore, _ := strconv.ParseInt(score, 10, 32) // 정수형 32비트 타입으로 형변환

	var grade string
	if realscore >= 60 {
		grade = ("A")
	} else {
		grade = ("BCDF")
	}
	fmt.Printf("%d점은 %s등급 입니다\n", realscore, grade)
}
