package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

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
			assertEqual(t, areEqual, test.areEqual)
		})
	}
}
