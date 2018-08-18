package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Tile rune

const (
	empty  Tile = ' '
	cross       = 'X'
	naught      = 'O'
)

func (t Tile) String() string {
	return fmt.Sprintf("%c", t)
}

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

var inputRegExp = regexp.MustCompile(`(\d+)\s*,?\s*(\d+)`)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)
	text, _ := reader.ReadString('\n')
	return text
}

func getXY() ([]int, error) {
	input := getInput("x, y")
	output := inputRegExp.FindStringSubmatch(input)
	if output == nil {
		return nil, fmt.Errorf("invalid x or y")
	}
	xy := make([]int, 2, 2)
	for i, digitStr := range output[1:] {
		digit, _ := strconv.Atoi(digitStr)
		xy[i] = digit
	}
	return xy, nil
}

func main() {
	for {
		b := NewBoard()
		players := []Tile{cross, naught}
		currentPlayerIndex := 0

		for i := 0; i < boardSize*boardSize; i++ {
			currentPlayer := players[currentPlayerIndex]
			fmt.Println(b)
			fmt.Printf("current player: %s\n", currentPlayer)
			for {
				xy, err := getXY()
				if err != nil {
					fmt.Printf("getting x and y failed: %s\n", err)
					continue
				}
				err = b.SetTile(currentPlayer, xy[0]-1, xy[1]-1)
				if err != nil {
					fmt.Printf("setting tile failed: %s\n", err)
					continue
				}
				break
			}
			currentPlayerIndex = (currentPlayerIndex + 1) % len(players)
		}
	}
}
