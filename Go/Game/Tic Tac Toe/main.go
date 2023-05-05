package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkWinOrDraw(board [3][3]string) string {
	// Check rows and columns for a win
	for i := 0; i < 3; i++ {
		if board[i][0] != " " && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
		if board[0][i] != " " && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}
	// Check diagonals for a win
	if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != " " && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}

	// Check for a draw
	draw := true
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row][col] == " " {
				draw = false
			}
		}
	}

	if draw {
		return "Draw"
	}

	return ""
}

func getUserInput(board [3][3]string, player string) (int, int) {
	var row, col int
	for {
		fmt.Printf("Enter the row and column (1-3) for player %s (e.g. 1,2): ", player)
		var input string
		fmt.Scanln(&input)
		coords := strings.Split(input, ",")
		if len(coords) != 2 {
			fmt.Println("Invalid input. Please enter the row and column separated by a comma.")
			continue
		}

		row, _ = strconv.Atoi(coords[0])
		col, _ = strconv.Atoi(coords[1])

		if row < 1 || row > 3 || col < 1 || col > 3 {
			fmt.Println("Invalid coordinates. Please enter values between 1 and 3.")
			continue
		}

		if board[row-1][col-1] != " " {
			fmt.Println("Cell already occupied. Please choose a different cell.")
			continue
		}

		break
	}

	return row - 1, col - 1
}

func displayBoard(board [3][3]string) {
	fmt.Println("-------------")
	for row := 0; row < 3; row++ {
		fmt.Print("|")
		for col := 0; col < 3; col++ {
			fmt.Printf(" %s |", board[row][col])
		}
		fmt.Println("\n-------------")
	}
}

func main() {
	var board [3][3]string

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			board[row][col] = " "
		}
	}

	// display the board
	displayBoard(board)

	var currentPlayer string
	var result string
	for {
		currentPlayer = "X"
		row, col := getUserInput(board, currentPlayer)
		board[row][col] = currentPlayer
		displayBoard(board)
		result = checkWinOrDraw(board)
		if result != "" {
			break
		}

		currentPlayer = "O"
		row, col = getUserInput(board, currentPlayer)
		board[row][col] = currentPlayer
		displayBoard(board)
		result = checkWinOrDraw(board)
		if result != "" {
			break
		}
	}

	if result == "Draw" {
		fmt.Println("It's a draw!")
	} else {
		fmt.Printf("Player %s wins!\n", result)
	}
}
