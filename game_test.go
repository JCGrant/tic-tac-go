package tic

import (
	"fmt"
	"testing"
)

func TestTilesAreSame(t *testing.T) {
	tests := []struct {
		tiles    []Tile
		areEqual bool
	}{
		{
			[]Tile{},
			true,
		},
		{
			[]Tile{cross},
			true,
		},
		{
			[]Tile{cross, cross},
			true,
		},
		{
			[]Tile{naught, naught, naught},
			true,
		},
		{
			[]Tile{naught, naught, empty},
			false,
		},
		{
			[]Tile{naught, cross, empty},
			false,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			areEqual := tilesAreSame(test.tiles...)
			assertEqual(t, test.areEqual, areEqual)
		})
	}
}
