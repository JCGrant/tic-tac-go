package main

import "fmt"

type Tile rune
type Board [][]Tile

const (
	empty  Tile = ' '
	cross       = 'X'
	naught      = 'O'
)

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

func (b *Board) String() (result string) {
	for _, row := range *b {
		for _, r := range row {
			result += string(r)
		}
		result += "\n"
	}
	return
}

func main() {
	b := NewBoard()
	fmt.Println(b)
}
