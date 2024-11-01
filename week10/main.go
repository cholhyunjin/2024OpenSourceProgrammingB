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
	//fmt.Printf("%f\n", math.Sqrt(25.0))
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
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 { // 2를 제외한 모든 짝수는 소수가 아님
		isPrime = false
	} else { // 3이상의 홀수
		i := 3
		//fot i < int(math.Sqrt(float64(n))) {
		for i*i <= n {
			if n%i == 0 {
				isPrime = false
				break // 1과 자기자신을 제외한 첫 번째 약수가 발견 되면 반복문 종료
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
