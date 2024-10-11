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
	score = (strings.TrimSpace(score))            //줄바꿈, 띄어쓰기, 탭 등 제거 (python strip과 유사)
	realscore, _ := strconv.ParseFloat(score, 64) // 실수형 64비트 타입으로 형변환

	if realscore >= 90 {
		fmt.Println("A")
	} else {
		fmt.Println("BCDF")
	}
}
