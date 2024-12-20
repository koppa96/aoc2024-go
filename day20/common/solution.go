package common

import (
	"bufio"
	"math"
	"os"
	"strings"
)

type TreeNode struct {
	Pos Pos
	Len int
}

func CountCheatsWithOver100Improvement(maxCheatLength int) int {
	file, err := os.Open("inputs/day20.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var maze [][]rune
	var start *Pos
	var end *Pos

	for scanner.Scan() {
		line := scanner.Text()
		if start == nil {
			if idx := strings.IndexRune(line, 'S'); idx != -1 {
				start = &Pos{Row: len(maze), Col: idx}
			}
		}

		if end == nil {
			if idx := strings.IndexRune(line, 'E'); idx != -1 {
				end = &Pos{Row: len(maze), Col: idx}
			}
		}

		maze = append(maze, []rune(line))
	}

	if start == nil || end == nil {
		panic("failed to locate the start or end position")
	}

	path := traverseBfs(maze, *start, *end)

	count := 0
	for i := 0; i < len(path)-1; i++ {
		for j := i + 1; j < len(path); j++ {
			dist := distance(path[j], path[i])
			if dist <= maxCheatLength {
				improvement := j - (i + dist)
				if improvement >= 100 {
					count++
				}
			}
		}
	}

	return count
}

func traverseBfs(maze [][]rune, start, end Pos) []Pos {
	visited := make(map[Pos]int)
	queue := []TreeNode{{Pos: start, Len: 0}}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		if _, ok := visited[first.Pos]; ok || maze[first.Pos.Row][first.Pos.Col] == '#' {
			continue
		}

		visited[first.Pos] = first.Len
		if first.Pos == end {
			result := make([]Pos, len(visited))
			for pos, l := range visited {
				result[l] = pos
			}

			return result
		}

		for _, dir := range directions {
			next, ok := dir(first.Pos, maze)
			if !ok {
				continue
			}

			queue = append(queue, TreeNode{Pos: next, Len: first.Len + 1})
		}
	}

	return nil
}

func distance(dest, start Pos) int {
	horizontalDistance := math.Abs(float64(dest.Row - start.Row))
	verticalDistance := math.Abs(float64(dest.Col - start.Col))
	totalDistance := int(horizontalDistance + verticalDistance)

	return totalDistance
}
