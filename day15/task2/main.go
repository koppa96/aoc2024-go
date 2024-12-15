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
	var posX, posY int
	for scanner.Scan() && scanner.Text() != "" {
		row := make([]rune, 0, len(scanner.Text())*2)
		for i, r := range scanner.Text() {
			switch r {
			case '#':
				row = append(row, '#', '#')
			case 'O':
				row = append(row, '[', ']')
			case '.':
				row = append(row, '.', '.')
			case '@':
				row = append(row, '@', '.')
				posX, posY = len(board), 2*i
			}
		}
		board = append(board, row)
	}

	for scanner.Scan() {
		for _, instruction := range scanner.Text() {
			posX, posY = move(board, posX, posY, common.Direction(instruction))
		}
	}

	var sum int
	for i, row := range board {
		for j, col := range row {
			if col == '[' {
				sum += 100*i + j
			}
		}
	}

	fmt.Println(sum)
}

func canMove(board [][]rune, posX, posY int, dir common.Direction) bool {
	nextX, nextY := dir.Next(posX, posY)
	if board[nextX][nextY] == ']' {
		success := canMove(board, nextX, nextY, dir)

		if dir == common.Up || dir == common.Down {
			success = success && canMove(board, nextX, nextY-1, dir)
		}

		return success
	}

	if board[nextX][nextY] == '[' {
		success := canMove(board, nextX, nextY, dir)

		if dir == common.Up || dir == common.Down {
			success = success && canMove(board, nextX, nextY+1, dir)
		}

		return success
	}

	return board[nextX][nextY] == '.'
}

func move(board [][]rune, posX, posY int, dir common.Direction) (int, int) {
	if !canMove(board, posX, posY, dir) {
		return posX, posY
	}

	nextX, nextY := dir.Next(posX, posY)
	if board[nextX][nextY] == ']' {
		move(board, nextX, nextY, dir)

		if dir == common.Up || dir == common.Down {
			move(board, nextX, nextY-1, dir)
		}
	}

	if board[nextX][nextY] == '[' {
		move(board, nextX, nextY, dir)

		if dir == common.Up || dir == common.Down {
			move(board, nextX, nextY+1, dir)
		}
	}

	if board[nextX][nextY] == '.' {
		board[nextX][nextY] = board[posX][posY]
		board[posX][posY] = '.'
		return nextX, nextY
	}

	return posX, posY
}
