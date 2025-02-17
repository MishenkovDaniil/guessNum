package main

import (
	"fmt"
	"math/rand"
)

const greetingMsg = "Привет, это игра guessNum, где тебе потребуется угадать загаданное мною число!"
const startMsg = "Уровень выбран! Чтобы начать просто сделайте первое предположение!"
const easyLevel = "легкий"
const mediumLevel = "средний"
const hardLevel = "сложный"
const easyLevelMax = 20
const mediumLevelMax = 50
const hardLevelMax = 100

func readNum(num *int) error {
	_, err := fmt.Scan(num)

	if err != nil {
		fmt.Println("Введите корректное число!")
	}
	return err
}

func generateNum() int {
	fmt.Printf("Выберите уровень сложности:\n%v\n%v\n%v\n", easyLevel, mediumLevel, hardLevel)
	var level string
	var num int
	for {
		_, err := fmt.Scan(&level)
		if err != nil {
			fmt.Println("Введите корректное значение!")
			continue
		}

		switch level {
		case easyLevel:
			num = rand.Intn(easyLevelMax)
		case mediumLevel:
			num = rand.Intn(mediumLevelMax)
		case hardLevel:
			num = rand.Intn(hardLevelMax)
		default:
			fmt.Println("Введите корректное значение!")
			continue
		}
		break
	}
	return num
}

func main() {
	fmt.Println(greetingMsg)
	myNum := generateNum()
	fmt.Println(startMsg)

	var guessedNum, tries int
	for {
		if readNum(&guessedNum) != nil {
			continue
		}
		tries++
		if myNum > guessedNum {
			fmt.Println("Слишком мало!")
		} else if myNum < guessedNum {
			fmt.Println("Слишком много!")
		} else {
			fmt.Printf("Поздравляем! Вы угадали число за %v попыток\n", tries)
			break
		}
	}
}
