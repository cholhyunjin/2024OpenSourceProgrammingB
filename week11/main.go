package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"week11/keyboard"
)
// input
// 10
// 19

// output
// 11 13 17 19

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		for j := 3; j*j <= n; j = j + 2 {
			if n%j == 0 {
				return false
			}
			//fmt.Printf("%d ", j)
		}
	}
	return true
}

func getInteger() (int, error) {
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func main() {
	fmt.Print("첫 번째 정수(시작 값) 입력 : ")
	n1, err := getInteger()
	n1, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("두 번째 정수(끝 값) 입력 : ")
	n2, err := getInteger()
	n2, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}
}
