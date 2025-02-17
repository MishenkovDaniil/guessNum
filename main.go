package main

import (
	"fmt"
	"math/rand"
)

const GreetingMsg = "Привет, это игра guessNum, где тебе потребуется угадать загаданное мною число. Чтобы начать просто сделайте первое предположение!"

func main() {
	fmt.Println(GreetingMsg)
	myNum := rand.Intn(100)
	i := 1
	for {
		var guessedNum int
		_, err := fmt.Scan(&guessedNum)
		if err != nil {
			fmt.Println("Введите корректное число!")
			// log.Fatal(err)
		}
		// if n > 1 {
		// 	continue
		// }
		// guessedNum, err := strconv.Atoi(msg)
		if myNum > guessedNum {
			fmt.Println("Слишком мало!")
			i++
		} else if myNum < guessedNum {
			fmt.Println("Слишком много!")
			i++
		} else {
			fmt.Printf("Поздравляем! Вы угадали число за %v попыток\n", i)
			break
		}
	}
}
