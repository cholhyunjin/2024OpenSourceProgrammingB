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
	fmt.Print("정수 입력 : ")
	i := bufio.NewReader(os.Stdin)
	number, err := i.ReadString('\n')

	number = (strings.TrimSpace(number))
	n, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}

	var isPrime bool = true
	if n <= 1 {
		isPrime = false
	} else {
		i := 2
		for i < n {
			if n%i == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d", i) // 반복 횟수 확인용 코드
			i++
		}
	}
	if isPrime {
		fmt.Printf("%d는 소수입니다.", n)
	} else {
		fmt.Printf("%d는 소수가 아닙니다.", n)
	}
}
