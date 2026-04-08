package main

import (
	"fmt"
	"gameOfLife/internal/helpers"
	"os"
	"os/exec"
	"time"
	"unicode"
)

var rows, cols int

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	generation := 0
	// y, x := 0, 0

	main := unicode.ToLower(helpers.MainMenu())

	if main == 'e' {

		matrix := helpers.ReadFile()

		helpers.OutputMatrix(matrix, generation)

		cellAlive := helpers.CountAlive(matrix)

		newMatrix := helpers.CopyMatrix(matrix)

		clear()

		for {

			helpers.KillOrRevive(matrix, newMatrix)

			generation++

			helpers.OutputMatrix(newMatrix, generation)

			time.Sleep(time.Second)

			clear()

			cellAlive = helpers.CountAlive(newMatrix)

			if !helpers.GameNoOver(cellAlive) || helpers.WorldStagnated(matrix, newMatrix) {
				break
			}

			matrix = helpers.CopyMatrix(newMatrix)

		}

		if !helpers.GameNoOver(cellAlive) {

			fmt.Println("\nAll cells are dead!\n")
			time.Sleep(time.Second)

		} else if helpers.WorldStagnated(matrix, newMatrix) {

			fmt.Println("\nWorld stagnated!\n")
			time.Sleep(time.Second)

		}

		helpers.GameOverScreen()

		time.Sleep(time.Second)

		clear()

	} else if main == 'q' {

		helpers.GameOverScreen()

		time.Sleep(3 * time.Second)

		clear()

	} else {

		fmt.Println("WRONG INPUT TRY AGAIN.\n")

		time.Sleep(2 * time.Second)

		clear()

	}

}
