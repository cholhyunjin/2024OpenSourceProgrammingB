package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else { 
		i := 3
		for i*i <= n {
			if n%i == 0 {
				return false
			}
			fmt.Printf("%d", i) 
			i = i + 2
		}
	}
}

func main() {
	//fmt.Printf("%f\n", math.Sqrt(25.0))
	fmt.Print("정수 입력 : ")
	i := bufio.NewReader(os.Stdin)
	number, err := i.ReadString('\n')
	i := (strings.TrimSpace(number))
	number, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}


	}
	if isPrime(n) {
		fmt.Printf("%d는 소수입니다.", n)
	} else {
		fmt.Printf("%d는 소수가 아닙니다.", n)
	}
}
