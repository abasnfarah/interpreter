package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/abasnfarah/monkey-interpreter/go/lexer"
	"github.com/abasnfarah/monkey-interpreter/go/token"
)

const (
	PROMPT       = ">> "
	HELP_MESSAGE = "help - show this help message\n" + "exit - exit the repl\n"
	HELP         = "help"
	EXIT         = "exit"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == EXIT {
			return
		}

		if line == HELP || line == "h" || line == "" {
			fmt.Fprint(out, HELP_MESSAGE)
			continue
		}

		l := lexer.New(line)

		fmt.Fprintf(out, "%+v\n", l)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
