package tic

import "fmt"

// Board is a grid of tiles.
type Board [][]Tile

const boardSize = 3

// NewBoard initialises a new board, filling it with empty tiles.
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

// MustGetTile fetchs a tile from a board.
// It will panic if the indexes are out of the boards bounds.
func (b *Board) MustGetTile(x, y int) Tile {
	return [][]Tile(*b)[y][x]
}

// MustSetTile will set a tile on a board.
// It will panic if the indexes are out of the boards bounds.
func (b *Board) MustSetTile(tile Tile, x int, y int) {
	[][]Tile(*b)[y][x] = tile
}

// SetTile will first check if the indexes are within the boards bounds, or
// if a tile has already been placed. Otherwise it will return an error.
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
