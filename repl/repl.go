package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Sa2Knight/maron/lexer"
	"github.com/Sa2Knight/maron/token"
)

// PROMPT REPLに毎行表示する文字列
const PROMPT = ">> "

// Start REPLを開始する
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
