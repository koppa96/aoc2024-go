package day15

type Direction rune

const (
	Right Direction = '>'
	Down  Direction = 'v'
	Left  Direction = '<'
	Up    Direction = '^'
)

func (d Direction) Next(posX, posY int) (int, int) {
	switch d {
	case Right:
		return posX, posY + 1
	case Down:
		return posX + 1, posY
	case Left:
		return posX, posY - 1
	case Up:
		return posX - 1, posY
	}

	panic("unknown direction")
}
