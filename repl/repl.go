package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/zscrub/golly/lexer"
	"github.com/zscrub/golly/token"
)


var lineNumberPrompt string = "1 | "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	i := 2
	for {
		fmt.Fprintf(out, lineNumberPrompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		l := scanner.Text()
		line := lexer.New(l)
		for tok := line.NextToken(); 
			tok.Type != token.EOF; 
			tok = line.NextToken() {
				fmt.Fprintf(out, "%+v\n", tok)
		}

		// set line number in prompt		
		lineNumberPrompt = string(i) + " | "
		i++
	}
}