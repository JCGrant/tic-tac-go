package tic

import "fmt"

type Tile rune

const (
	empty  Tile = ' '
	cross       = 'X'
	naught      = 'O'
)

func (t Tile) String() string {
	return fmt.Sprintf("%c", t)
}
