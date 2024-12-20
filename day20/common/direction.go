package common

type Direction func(pos Pos, maze [][]rune) (Pos, bool)

var directions = [...]Direction{
	// Right
	func(pos Pos, maze [][]rune) (Pos, bool) {
		if pos.Col == len(maze[pos.Row])-1 {
			return Pos{}, false
		}

		return Pos{Row: pos.Row, Col: pos.Col + 1}, true
	},
	// Down
	func(pos Pos, maze [][]rune) (Pos, bool) {
		if pos.Row == len(maze)-1 {
			return Pos{}, false
		}

		return Pos{Row: pos.Row + 1, Col: pos.Col}, true
	},
	// Left
	func(pos Pos, maze [][]rune) (Pos, bool) {
		if pos.Col == 0 {
			return Pos{}, false
		}

		return Pos{Row: pos.Row, Col: pos.Col - 1}, true
	},
	// Up
	func(pos Pos, maze [][]rune) (Pos, bool) {
		if pos.Row == 0 {
			return Pos{}, false
		}

		return Pos{Row: pos.Row - 1, Col: pos.Col}, true
	},
}
