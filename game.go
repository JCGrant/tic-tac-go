package tic

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
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

func tilesAreSame(tiles ...Tile) bool {
	allEqual := true
	for i := 0; i < len(tiles)-1; i++ {
		currentTile := tiles[i]
		nextTile := tiles[i+1]
		allEqual = allEqual && (currentTile == nextTile)
	}
	return allEqual
}

func isWinningTileSet(tiles ...Tile) bool {
	if len(tiles) != boardSize {
		return false
	}
	if !tilesAreSame(tiles...) {
		return false
	}
	if tiles[0] == empty {
		return false
	}
	return true
}

func containsWinner(board *Board) (bool, Tile) {
	// will collect the diagonals while collecting the tile rows/columns
	forwardDiagonal := []Tile{}
	backwardDiagonal := []Tile{}
	for i := 0; i < boardSize; i++ {
		// Can collect the tile rows and columns in the same pass
		// because the board width == height.
		// Otherwise this would cause an index out of range panic
		row := []Tile{}
		column := []Tile{}
		for j := 0; j < boardSize; j++ {
			rowTile := board.MustGetTile(j, i)
			columnTile := board.MustGetTile(i, j)
			row = append(row, rowTile)
			column = append(column, columnTile)
			if i == j {
				forwardDiagonal = append(forwardDiagonal, rowTile)
			}
			if i+boardSize-1 == j {
				backwardDiagonal = append(backwardDiagonal, rowTile)
			}
		}
		if isWinningTileSet(row...) {
			return true, row[0]
		}
		if isWinningTileSet(column...) {
			return true, column[0]
		}
	}
	if isWinningTileSet(forwardDiagonal...) {
		return true, forwardDiagonal[0]
	}
	if isWinningTileSet(backwardDiagonal...) {
		return true, backwardDiagonal[0]
	}
	return false, empty
}

func (g *Game) Run() {
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
		if ok, tile := containsWinner(b); ok {
			fmt.Printf(`
#################################
        Player %s has won!
#################################`, tile)
			fmt.Println()
			break
		}
		currentPlayerIndex = (currentPlayerIndex + 1) % len(players)
	}
	fmt.Printf(`
#################################
          It's a draw!
#################################`)
	fmt.Println()
}
