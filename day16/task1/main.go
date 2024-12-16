package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	East Direction = iota
	South
	West
	North
)

func (d Direction) String() string {
	switch d {
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	case North:
		return "North"
	}

	return fmt.Sprintf("Unknown (%d)", d)
}

type Pos struct {
	Row int
	Col int
	Dir Direction
}

func (pos Pos) Next() Pos {
	switch pos.Dir {
	case North:
		return Pos{Row: pos.Row - 1, Col: pos.Col, Dir: pos.Dir}
	case East:
		return Pos{Row: pos.Row, Col: pos.Col + 1, Dir: pos.Dir}
	case South:
		return Pos{Row: pos.Row + 1, Col: pos.Col, Dir: pos.Dir}
	case West:
		return Pos{Row: pos.Row, Col: pos.Col - 1, Dir: pos.Dir}
	}

	panic(fmt.Errorf("unknown direction: %d", pos.Dir))
}

func (pos Pos) RotateClockwise() Pos {
	return Pos{Row: pos.Row, Col: pos.Col, Dir: (pos.Dir + 1) % 4}
}

func (pos Pos) RotateCounterclockwise() Pos {
	if pos.Dir == East {
		return Pos{Row: pos.Row, Col: pos.Col, Dir: North}
	}

	return Pos{Row: pos.Row, Col: pos.Col, Dir: pos.Dir - 1}
}

type State struct {
	Pos    Pos
	Points int64
}

func main() {
	file, err := os.Open("inputs/day16.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var board [][]rune
	for scanner.Scan() {
		board = append(board, []rune(scanner.Text()))
	}

	start, ok := findPos(board, 'S')
	if !ok {
		panic("failed to find starting position")
	}

	end, ok := findPos(board, 'E')
	if !ok {
		panic("failed to start ending position")
	}

	start.Dir = East
	minPoints := traverseOptimal(board, start, end)

	fmt.Println(minPoints)
}

func findPos(board [][]rune, typ rune) (Pos, bool) {
	for row := range board {
		for col := range board[row] {
			if board[row][col] == typ {
				return Pos{Row: row, Col: col}, true
			}
		}
	}

	return Pos{}, false
}

func traverseOptimal(board [][]rune, start, end Pos) int64 {
	minPoints := int64(-1)
	visited := map[Pos]int64{
		start: 0,
	}

	queue := []State{
		{Pos: start.RotateClockwise(), Points: 1000},
		{Pos: start.RotateCounterclockwise(), Points: 1000},
	}

	if next := start.Next(); board[next.Row][next.Col] != '#' {
		queue = append(queue, State{Pos: next, Points: 1})
	}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		if first.Pos.Row == end.Row && first.Pos.Col == end.Col {
			if minPoints == -1 || first.Points < minPoints {
				minPoints = first.Points
			}

			continue
		}

		next := first.Pos.Next()
		if board[next.Row][next.Col] != '#' {
			if points, ok := visited[next]; !ok || points > first.Points+1 {
				visited[next] = first.Points + 1
				queue = append(queue, State{Pos: next, Points: first.Points + 1})
			}
		}

		next = first.Pos.RotateClockwise()
		if points, ok := visited[next]; !ok || points > first.Points+1000 {
			visited[next] = first.Points + 1000
			queue = append(queue, State{Pos: next, Points: first.Points + 1000})
		}

		next = first.Pos.RotateCounterclockwise()
		if points, ok := visited[next]; !ok || points > first.Points+1000 {
			visited[next] = first.Points + 1000
			queue = append(queue, State{Pos: next, Points: first.Points + 1000})
		}
	}

	return minPoints
}
