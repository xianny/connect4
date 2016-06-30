package main

import (
	"bufio"
	"github.com/jwils/connect4/game"
	"github.com/jwils/connect4/players"
	"os"
)

func main() {
	p1 := players.RandomPlayer{}
	p2 := players.HumanPlayer{}
	p2.Reader = bufio.NewReader(os.Stdin)
	game := game.NewGame(p1, p2)
	game.Play()
}
