package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
	var guess int
	ret := 0
	for ret != 1 {
		fmt.Println("Please input your guess\n")
		// 输入我们猜的数字
		fmt.Scanln()
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value\n")
			return
		}
		fmt.Printf("Your guess is %d\r\n", guess)

		if guess < secretNumber {
			fmt.Println("Your guess is lower")
		} else if guess > secretNumber {
			fmt.Println("Your guess is higher")
		} else {
			fmt.Println("Yes,you are right!")
			ret = 1
		}
	}
	fmt.Println("Game is over")
}
