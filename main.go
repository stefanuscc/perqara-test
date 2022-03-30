package main

import (
	"fmt"
	"sort"
)

/** Fraud Notifications */

func median(list []int) float32 {
	// Sort the list
	sort.Ints(list)

	// If the length of the array is odd
	if len(list)%2 == 1 {
		return float32(list[len(list)/2])
	} else {
		return float32((list[len(list)/2] + list[len(list)/2-1])) / 2
	}
}

func fraudNotif(d int, expenditure []int) int {
	// Init counter
	count := 0

	// Iterate from d to len(expenditure)
	for i := d; i < len(expenditure); i++ {
		// Get the median
		median := median(expenditure[i-d : i])
		// If the current expenditure is greater than the median
		if float32(expenditure[i]) >= median*2 {
			// Increment the counter
			count++
		}
	}

	// Return Count
	return count
}

/** Queen's Attack */
type Chessboard struct {
	boardSize      int
	obstaclesCount int
	queenX         int
	queenY         int
	obstacles      [][]int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isCellAttackable(x int, y int, board *Chessboard) bool {
	// Out of board check
	if x > board.boardSize || y > board.boardSize {
		return false
	}

	// Orthogonal check
	if x == board.queenX || y == board.queenY {
		return true
	}

	// Diagonal check
	if abs(x-board.queenX) == abs(y-board.queenY) {
		return true
	}

	return false
}

func isCellBlocked(x int, y int, board *Chessboard) bool {
	// Out of board check
	if x > board.boardSize || y > board.boardSize {
		return true
	}

	// Iterate obstacles
	for _, obs := range board.obstacles {
		if (obs[0]%x == 0) && (obs[1]%y == 0) {
			return true
		}
	}

	return false
}

func queensAttack(boardSize int, obstaclesCount int, queenY int, queenX int, obstacles [][]int) int {
	// Init Chessboard
	var chessboard Chessboard
	chessboard.boardSize = boardSize
	chessboard.obstaclesCount = obstaclesCount
	chessboard.queenX = queenX
	chessboard.queenY = queenY
	chessboard.obstacles = obstacles

	// Init stuck stats
	// stuckUl := false
	// stuckUc := false
	// stuckUr := false
	// stuckCl := false
	// stuckCr := false
	// stuckBl := false
	// stuckBc := false
	// stuckBr := false
	counter := 0

	// Iterate board row
	for y := 1; y <= boardSize; y++ {
		// Iterate board columns
		for x := 1; x <= boardSize; x++ {
			if isCellAttackable(x, y, &chessboard) {
				if !isCellBlocked(x, y, &chessboard) {
					fmt.Println(x, y)
					counter++
				}
			}
		}
	}

	return counter
}

func main() {
	fmt.Println("Hello Perqara, I'm Stefan")
	fmt.Println("Fraud Notification Test")
	fmt.Println("Fraud Test Case 1: ", fraudNotif(3, []int{10, 20, 30, 40, 50}))
	fmt.Println("Fraud Test Case 2: ", fraudNotif(5, []int{2, 3, 4, 2, 3, 6, 8, 4, 5}))
	fmt.Println("Fraud Test Case 3: ", fraudNotif(4, []int{1, 2, 3, 4, 4}))

	fmt.Println("Queen's Attack")
	fmt.Println("Queen Test Case 1: ", queensAttack(8, 1, 4, 4, [][]int{{3, 5}}))
	fmt.Println("Queen Test Case 2: ", queensAttack(5, 3, 4, 3, [][]int{{5, 5}, {4, 2}, {2, 3}}))
}
