/*
A simple ascii x_o game

1. Define a function to output game state

*/

package main

import "fmt"

// Check if someone won the game
func game_done(game []string) (bool, string) {
	winning_combinations := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	for _, value := range winning_combinations {
		if game[value[0]] != "" && game[value[0]] == game[value[1]] && game[value[1]] == game[value[2]] {
			return true, game[value[0]]
		}
	}
	return false, ""
}

// Check if game is a draw
func game_draw(game []string) bool {
	for _, value := range game {
		if value == "" {
			return false
		}
	}
	return true
}

// Display game board
func display(game []string) {
	for i := 0; i < len(game); i += 3 {
		fmt.Printf("| %s | %s | %s |\n", game[i], game[i+1], game[i+2])
		fmt.Println("---------")
	}
}

func get_input(game []string) (string, int) {
	// Gets user input
	// Returns user input
	// https://www.w3schools.com/go/go_input.php

	// Requires the pointer as it modifies the actual value.
	// Simple var passing returns a copy

	for {
		// Get player side
		var side string
		fmt.Println("Please enter side, X or O")
		fmt.Scan(&side)
		if side != "X" && side != "O" {
			fmt.Println("Invalid side")
			continue
		}

		// Get user position
		var position int
		fmt.Println("Please enter position, 0-9")
		fmt.Scan(&position)
		if game[position] != "" {
			fmt.Println("Invalid position, square is already taken")
			continue
		}
		if (position > 9) || (position < 0) {
			fmt.Println("Invalid position, out of range")
			continue
		}
		return side, position
	}
}

func main() {

	// Initialize game state
	game := make([]string, 9)
	var run_game bool = true

	// While game is not finished, run
	for run_game {

		// Check if there is a winner
		done, winner := game_done(game)
		if done {
			fmt.Printf("User %s won\n", winner)
			break
		}

		// Check for a draw
		draw := game_draw(game)
		if draw {
			fmt.Println("Game is a draw")
			break
		}

		// Display current board
		display(game)

		// Get user input
		side, position := get_input(game)

		// Update board
		game[position] = side

		// Display updated board
		display(game)

	}
	fmt.Println("Game Over")
}
