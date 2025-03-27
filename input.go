package main

import (
	"fmt"
)

func input(type_game int) {
	var i int
	var k int
	fmt.Println("Введите координаты ячейки которую хотите открыть.")
	for {
		fmt.Printf("Введите номер строки:")
		fmt.Scan(&i)
		fmt.Printf("Введите номер столбца:")
		fmt.Scan(&k)

		if i-1 < 0 || i-1 >= type_game || k-1 < 0 || k-1 >= type_game {
			fmt.Println("Вы ввели не верные координаты. Попробуйте ещё раз:")
			continue
		}

		if game_board[i-1][k-1] != "[-]" {
			fmt.Println("Вы ввели координаты уже открытой ячейки. Попробуйте ещё раз:")
		} else {
			break
		}
	}

	check_progress(i-1, k-1, type_game)
}
