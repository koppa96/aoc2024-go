package main

import (
	"bufio"
	"fmt"
	"github.com/koppa96/aoc2024go/day15/common"
	"os"
)

func main() {
	file, err := os.Open("inputs/day15.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var board [][]rune
	for scanner.Scan() && scanner.Text() != "" {
		board = append(board, []rune(scanner.Text()))
	}

	posX, posY, ok := findRobotPos(board)
	if !ok {
		panic("failed to find the robot on the board")
	}

	for scanner.Scan() {
		for _, instruction := range []rune(scanner.Text()) {
			posX, posY = move(board, posX, posY, common.Direction(instruction))
		}
	}

	var sum int
	for i, row := range board {
		for j, col := range row {
			if col == 'O' {
				sum += 100*i + j
			}
		}
	}

	fmt.Println(sum)
}

func findRobotPos(board [][]rune) (int, int, bool) {
	for i, row := range board {
		for j, col := range row {
			if col == '@' {
				return i, j, true
			}
		}
	}

	return 0, 0, false
}

func move(board [][]rune, posX, posY int, dir common.Direction) (int, int) {
	nextX, nextY := dir.Next(posX, posY)
	if board[nextX][nextY] == 'O' {
		move(board, nextX, nextY, dir)
	}

	if board[nextX][nextY] == '.' {
		board[nextX][nextY] = board[posX][posY]
		board[posX][posY] = '.'
		return nextX, nextY
	}

	return posX, posY
}
