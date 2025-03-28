package main

import (
	"fmt"
)

var count_bombs int
var game_over bool
var game_board [][]string
var board_map [][]string

func game(type_game int) {
	var expectation string
	board_map_start(type_game)
	game_board = make([][]string, type_game)
	for i := range game_board {
		game_board[i] = make([]string, type_game)
	}
	board_start(type_game)

	for !game_over {
		input(type_game)
		clear()
		board(type_game)
	}

	if count_bombs == type_game {
		fmt.Println("Вы победили!")
		fmt.Println("\nВведите любой символ чтобы начать игру заново.")
		fmt.Scan(&expectation)
		main()
	} else if game_over == true {
		fmt.Println("Вы проиграли")
		fmt.Println("\nВведите любой символ чтобы начать игру заново.")
		fmt.Scan(&expectation)
		main()
	}
}
