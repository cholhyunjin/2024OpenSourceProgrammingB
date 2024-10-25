package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	fmt.Println(answer)

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Print("숫자 입력 : ")
		i := bufio.NewReader(os.Stdin)
		input, err := i.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input) //줄바꿈, 띄어쓰기, 탭 등 제거 (python strip과 유사)
		// guess, err := strconv.ParseInt(input, 10, 32)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답이에요")
			break
		} else if answer > guess {
			fmt.Println("입력하신 값은 정답보다 작은 값입니다.")
		} else {
			fmt.Println("입력 하신 값은 정답보다 큰 값입니다")
		}
	}
}
