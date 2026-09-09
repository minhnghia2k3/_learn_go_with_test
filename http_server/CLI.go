package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	reader      *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		reader:      bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()

	cli.playerStore.RecordWin(extractWinner(userInput))
}

// Chris win
func extractWinner(text string) string {
	return strings.Replace(text, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.reader.Scan()
	return cli.reader.Text()
}
