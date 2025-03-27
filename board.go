package main

import (
	"fmt"
)

func board_map_start(type_board int) {
	board_map = make([][]string, type_board)
	for i := range board_map {
		board_map[i] = make([]string, type_board)
	}

	for i := 0; i < type_board; i++ {
		j := random(type_board)
		k := random(type_board)
		if board_map[j][k] != "[X]" {
			board_map[j][k] = "[X]"
		}
	}

	for i := 0; i < type_board; i++ {
		for k := 0; k < type_board; k++ {
			if board_map[i][k] != "[X]" {
				count := bomb_around(i, k, type_board)
				if count != 0 {
					board_map[i][k] = fmt.Sprintf("[%d]", count)
				} else {
					board_map[i][k] = "[ ]"
				}
			}
		}
	}
}

func bomb_around(i int, k int, type_board int) int {
	count := 0
	directions := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		nx, ny := i+dir.dx, k+dir.dy
		if nx >= 0 && ny >= 0 && nx < type_board && ny < type_board {
			if board_map[nx][ny] == "[X]" {
				count += 1
			}
		}
	}
	return count

}

func board_start(type_board int) {
	for i := 0; i < type_board; i++ {
		for k := 0; k < type_board; k++ {
			game_board[i][k] = "[-]"
		}
	}
	board(type_board)
}

func board(type_board int) {
	for i := 0; i < type_board; i++ {
		for k := 0; k < type_board; k++ {
			fmt.Printf("%s", game_board[i][k])
		}
		fmt.Printf("\n")
	}
}
