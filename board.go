package tic

import "fmt"

type Board [][]Tile

const boardSize = 3

func NewBoard() *Board {
	b := Board{}
	for i := 0; i < boardSize; i++ {
		var row []Tile
		for j := 0; j < boardSize; j++ {
			row = append(row, empty)
		}
		b = append(b, row)
	}
	return &b
}

func horizontalGridLine() (result string) {
	result += "+"
	for i := 0; i < boardSize; i++ {
		result += " - +"
	}
	result += "\n"
	return
}

func (b *Board) String() (result string) {
	result += "    "
	for i := 0; i < boardSize; i++ {
		result += fmt.Sprintf("%d   ", i+1)
	}
	result += "\n"
	result += "  "
	result += horizontalGridLine()
	for i, row := range *b {
		result += fmt.Sprintf("%d ", i+1)
		result += "|"
		for _, t := range row {
			result += fmt.Sprintf(" %s |", t)
		}
		result += "\n"
		result += "  "
		result += horizontalGridLine()
	}
	return
}

func (b *Board) MustGetTile(x, y int) Tile {
	return [][]Tile(*b)[y][x]
}

func (b *Board) MustSetTile(tile Tile, x int, y int) {
	[][]Tile(*b)[y][x] = tile
}

func (b *Board) SetTile(tile Tile, x int, y int) error {
	if x < 0 || x >= boardSize {
		return fmt.Errorf("x is out of bounds")
	}
	if y < 0 || y >= boardSize {
		return fmt.Errorf("y is out of bounds")
	}
	if b.MustGetTile(x, y) != empty {
		return fmt.Errorf("a tile has already been placed at %d, %d", x, y)
	}
	b.MustSetTile(tile, x, y)
	return nil
}
