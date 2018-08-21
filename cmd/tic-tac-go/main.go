package main

import tic "github.com/JCGrant/tic-tac-go"

func main() {
	for {
		g := tic.NewGame()
		g.Run()
	}
}
