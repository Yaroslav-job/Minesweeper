package main

func check_progress(i int, k int, type_board int) {
	if board_map[i][k] == "[X]" {
		game_board[i][k] = board_map[i][k]
		game_over = true
	} else if board_map[i][k] != "[ ]" && board_map[i][k] != "[X]" {
		game_board[i][k] = board_map[i][k]
	} else if game_board[i][k] != "[-]" {
		return
	} else if board_map[i][k] == "[ ]" {
		game_board[i][k] = board_map[i][k]

		directions := []struct{ dx, dy int }{
			{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
		}

		for _, dir := range directions {
			nx, ny := i+dir.dx, k+dir.dy
			if nx >= 0 && ny >= 0 && nx < type_board && ny < type_board {
				check_progress(nx, ny, type_board)
			}
		}
	}
	check_win(type_board)
}

func check_win(type_board int) {
	count_bombs = 0
	for i := 0; i < type_board; i++ {
		for k := 0; k < type_board; k++ {
			if game_board[i][k] == "[-]" {
				count_bombs += 1
			}
		}
	}

	if count_bombs == type_board {
		game_over = true
	}
}
