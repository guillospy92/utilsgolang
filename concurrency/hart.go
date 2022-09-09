package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var containerGame = map[int][]string{
	1: {" ", " ", " "},
	2: {" ", "x", " "},
	3: {" ", " ", " "},
}

var limitBox = 8

func main() {
	found := func(x []int32, numEvaluated int32) bool {
		var count int32
		for _, n := range x {
			if n == numEvaluated {
				count += 1
			}
		}

		if count > 1 {
			return true
		}

		return false
	}

	fmt.Println(a(2))

	fmt.Println("tic-tac-toe game")
	fmt.Println("Write your next move and press enter: 1,1 example")
	fmt.Println("---------------------")

	for {
		paintContainer()
		fmt.Print("->")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		positionContainer, positionBox, err := validateGame(scanner.Text())

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// player
		containerGame[positionContainer][positionBox] = "0"
		limitBox -= 1
		if verifyEndGame(positionContainer, positionBox) {
			paintContainer()
			fmt.Println("winner")
			return
		}

		// machine
		positionContainer, positionBox = radonMachine()
		limitBox -= 1
		if verifyEndGame(positionContainer, positionBox) {
			paintContainer()
			fmt.Println("game over loser")
			return
		}
		// 1,1, 2,3, 2,2, 3,2, 1,1
		fmt.Println(limitBox, "Limite")
		if limitBox == 0 {
			fmt.Println("game over nobody wins")
			paintContainer()
			return
		}
	}
}

func radonMachine() (int, int) {
	for i := 1; i <= 3; i++ {
		for j := 0; j <= 2; j++ {
			if containerGame[i][j] == " " {
				containerGame[i][j] = "x"
				return i, j
			}
		}
	}
	return 0, 0
}

func verifyEndGame(positionContainer int, positionBox int) bool {
	// search vertical
	if searchVertical(positionBox) {
		return true
	}

	if searchHorizontal(positionContainer) {
		return true
	}

	return searchDiagonal(positionContainer)
}

func validateGame(movementText string) (positionContainer int, positionBox int, err error) {

	movement := strings.Split(movementText, ",")

	if len(movement) != 2 {
		return 0, 0, errors.New("incorrect patter, example 1,1")
	}

	positionContainer, err = strconv.Atoi(movement[0])
	positionBox, err = strconv.Atoi(movement[1])

	if err != nil {
		return 0, 0, errors.New("incorrect patter only number")
	}

	if positionContainer > 3 || positionBox > 2 {
		return 0, 0, errors.New("incorrect patter first position max 3 and second position max 2")
	}

	if containerGame[positionContainer][positionBox] != " " {
		return 0, 0, errors.New("the box is busy")
	}

	return positionContainer, positionBox, nil
}

func paintContainer() {
	fmt.Println("+---+---+---+")
	for i := 1; i <= 3; i++ {
		fmt.Printf("| %s | %s | %s | \n", containerGame[i][0], containerGame[i][1], containerGame[i][2])
		fmt.Println("+---+---+---+")
	}
}

func searchVertical(positionBox int) bool {
	first := containerGame[1][positionBox]
	second := containerGame[2][positionBox]
	three := containerGame[3][positionBox]

	return first == second && first == three
}

func searchHorizontal(positionContainer int) bool {
	first := containerGame[positionContainer][0]
	second := containerGame[positionContainer][1]
	three := containerGame[positionContainer][2]
	return first == second && first == three
}

func searchDiagonal(positionContainer int) bool {
	// search position down
	if positionContainer == 3 {
		first := containerGame[1][0]
		second := containerGame[2][1]
		three := containerGame[3][2]
		return first == second && first == three
	}

	// search position up
	if positionContainer == 1 {
		first := containerGame[1][2]
		second := containerGame[2][1]
		three := containerGame[3][0]

		return first == second && first == three
	}

	return false
}
