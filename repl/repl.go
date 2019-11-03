package repl

import (
	"bufio"
	"fmt"
	"io"

	"prashant.com/monkey/lexer"
	"prashant.com/monkey/token"
)

// PROMPT constant
const PROMPT = ">> "

// Start the repl
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
