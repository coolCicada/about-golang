package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println(secretNumber)

	fmt.Println("please input you guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error")
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		
		if guess > secretNumber {
			fmt.Println("bigger")
		} else if guess < secretNumber {
			fmt.Println("smaller")
		} else {
			fmt.Println("Success!!!!")
			break
		}
	}
}