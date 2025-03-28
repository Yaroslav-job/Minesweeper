package main

import (
	"fmt"
	"os"
)

func main() {
	var option string

	game_over = false

	clear()

	fmt.Println("Добро пожаловать в игру \"Сапёр\"")
	fmt.Println("Выберите сложность игры: \n1 - Лёгкий \n2 - Среднй \n3 - Сложный \nx - Закрыть программу")
	for {
		fmt.Scan(&option)
		if option == "1" {
			game(7)
			break
		} else if option == "2" {
			game(9)
			break
		} else if option == "3" {
			game(12)
			break
		} else if option == "x" {
			os.Exit(0)
		} else {
			fmt.Println("Вы ввели недопустимое значение. Попробуйте ещё раз:")
		}
	}
}
