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

	counts := 0
	i := 1
	for i <= n {
		if n%i == 0 {
			counts = counts + 1
		}
		i++
	}
	if counts == 2 {
		fmt.Printf("%d는 소수입니다.", n)
	} else {
		fmt.Printf("%d는 소수가 아닙니다.", n)
	}
}
