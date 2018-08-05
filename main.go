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
	result += horizontalGridLine()
	for _, row := range *b {
		result += "|"
		for _, t := range row {
			result += fmt.Sprintf(" %s |", t)
		}
		result += "\n"
		result += horizontalGridLine()
	}
	return
}

func (b *Board) MustSetTile(tile Tile, x int, y int) {
	[][]Tile(*b)[y][x] = tile
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
	b := NewBoard()
	fmt.Println(b)
	xy, err := getXY()
	if err != nil {
		fmt.Printf("getting x and y failed: %s", err)
	}
	b.MustSetTile(cross, xy[0], xy[1])
	fmt.Println(b)
}
