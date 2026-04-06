package helpers

import (
	"fmt"
	"os"
)

const CELL = 'O'
const DEAD_CELL = '.'

var rows, cols int

func MainMenu() rune {
	fmt.Println("   ***   Universe size and index of cell written in conf.  ***")
	fmt.Println(`
 ▄▄ •  ▄▄▄· • ▌ ▄ ·. ▄▄▄ .          ·▄▄▄    ▄▄▌  ▪  ·▄▄▄▄▄▄ .
▐█ ▀ ▪▐█ ▀█ ·██ ▐███▪▀▄.▀·    ▪     ▐▄▄·    ██•  ██ ▐▄▄·▀▄.▀·
▄█ ▀█▄▄█▀▀█ ▐█ ▌▐▌▐█·▐▀▀▪▄     ▄█▀▄ ██▪     ██▪  ▐█·██▪ ▐▀▀▪▄
▐█▄▪▐█▐█ ▪▐▌██ ██▌▐█▌▐█▄▄▌    ▐█▌.▐▌██▌.    ▐█▌▐▌▐█▌██▌.▐█▄▄▌
·▀▀▀▀  ▀  ▀ ▀▀  █▪▀▀▀ ▀▀▀      ▀█▄▀▪▀▀▀     .▀▀▀ ▀▀▀▀▀▀  ▀▀▀
`)
	fmt.Println("Press E to start")
	fmt.Println("Press Q to quit")

	var choice string
	fmt.Scan(&choice)

	return rune(choice[0])
}

func GameOverScreen() {
	fmt.Println(`
    ▄████  ▄▄▄       ███▄ ▄███▓▓█████     ▒█████   ██▒   █▓▓█████  ██▀███
   ██▒ ▀█▒▒████▄    ▓██▒▀█▀ ██▒▓█   ▀    ▒██▒  ██▒▓██░   █▒▓█   ▀ ▓██ ▒ ██▒
  ▒██░▄▄▄░▒██  ▀█▄  ▓██    ▓██░▒███      ▒██░  ██▒ ▓██  █▒░▒███   ▓██ ░▄█ ▒
  ░▓█  ██▓░██▄▄▄▄██ ▒██    ▒██ ▒▓█  ▄    ▒██   ██░  ▒██ █░░▒▓█  ▄ ▒██▀▀█▄
  ░▒▓███▀▒ ▓█   ▓██▒▒██▒   ░██▒░▒████▒   ░ ████▓▒░   ▒▀█░  ░▒████▒░██▓ ▒██▒
`)
}

func ReadFile() [][]byte {

	file, err := os.Open("./config/game_config.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fscan(file, &rows, &cols)

	matrix := make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
	}

	var i, j int

	for {
		_, err := fmt.Fscan(file, &i, &j)
		if err != nil {
			break
		}

		if i < rows && j < cols {
			matrix[i][j] = CELL
		}
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if matrix[y][x] != CELL {
				matrix[y][x] = DEAD_CELL
			}
		}
	}

	return matrix
}

func isAlive(matrix [][]byte, x, y int) int {
	if matrix[x][y] == CELL {
		return 1
	}
	return 0
}

func CountAlive(matrix [][]byte) int {

	cellAlive := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if isAlive(matrix, i, j) == 1 {
				cellAlive++
			}
		}
	}

	return cellAlive
}

func CopyMatrix(matrix [][]byte) [][]byte {

	newMatrix := make([][]byte, rows)

	for i := range newMatrix {
		newMatrix[i] = make([]byte, cols)
		copy(newMatrix[i], matrix[i])
	}

	return newMatrix
}

func OutputMatrix(matrix [][]byte, generation int) {

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%c", matrix[i][j])
		}
		fmt.Println()
	}

	cellAlive := CountAlive(matrix)

	fmt.Println("Generation:", generation)
	fmt.Println("Alive cells:", cellAlive)
}

func GameNoOver(cellAlive int) bool {
	return cellAlive != 0
}

func WorldStagnated(matrix, newMatrix [][]byte) bool {

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] != newMatrix[i][j] {
				return false
			}
		}
	}

	return true
}

func CountNeighbours(matrix [][]byte, y, x int) int {

	count := 0

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {

			if matrix[(i+rows)%rows][(j+cols)%cols] == CELL {
				count++
			}

		}
	}

	if matrix[y][x] == CELL {
		count--
	}

	return count
}

func KillOrRevive(matrix, newMatrix [][]byte) {

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			neighbours := CountNeighbours(matrix, i, j)

			if neighbours < 2 {
				newMatrix[i][j] = DEAD_CELL
			} else if neighbours > 3 {
				newMatrix[i][j] = DEAD_CELL
			} else if neighbours == 3 {
				newMatrix[i][j] = CELL
			} else {
				newMatrix[i][j] = matrix[i][j]
			}

		}
	}

}
