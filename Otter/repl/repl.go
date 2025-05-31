package repl

import (
	"bufio"
	"fmt"

	"io"

	"github.com/charmbracelet/lipgloss"
	"github.com/r-priyanshu/interpreter/evaluator"
	"github.com/r-priyanshu/interpreter/lexer"
	"github.com/r-priyanshu/interpreter/parser"
)

const OTTER_FACE = `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⣤⣤⣤⣴⣶⣶⣶⣦⣤⣤⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⡾⠿⠛⠛⠋⠉⠉⠉⠉⠉⠉⡉⠉⠉⠉⠛⠛⠿⢷⣶⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⡿⠟⠋⠁⠀⢀⡀⠀⠀⢰⡆⠀⠀⠀⢸⡇⠀⠀⠀⣀⠀⠀⠀⠈⠙⠻⢿⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⡿⠛⠁⠀⠀⠀⠀⠀⢸⣿⠀⠀⢸⡇⠀⠀⠀⢸⡇⠀⠀⠀⣿⡆⠀⠀⠀⠀⠀⠀⠉⠻⢷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢀⣠⣴⣶⠟⠋⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⠀⠀⢸⡇⠀⠀⠀⢸⡇⠀⠀⠀⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠿⣶⣦⣄⡀⠀⠀⠀⠀
⠀⠀⢠⣾⠿⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠛⠀⠀⠸⡧⠀⠀⠀⠘⠃⠀⠀⠀⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⢿⣷⡄⠀⠀
⠀⠀⢻⣿⣶⣶⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⢀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣶⣶⣿⠏⠀⠀
⠀⠀⠀⠙⠿⣷⣾⡇⠀⠀⠀⠀⠀⠀⢀⣤⣤⣄⠀⠀⠀⠀⠀⠀⠀⠉⠘⠃⠀⠀⠀⠀⠀⠀⢀⣤⣦⣤⡀⠀⠀⠀⠀⠀⠀⢻⣷⣾⠟⠋⠀⠀⠀
⠀⠀⠀⠀⠀⣸⡿⠁⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⡧⠀⠀⠀⠰⣶⣶⣾⣾⣷⣶⣶⠄⠀⠀⠀⣿⣿⣿⣿⡷⠀⠀⠀⠀⠀⠀⠈⣿⣇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣰⡿⠁⠀⠀⠀⠀⠀⠀⠀⠙⠻⠟⠛⠁⠀⠀⠀⠀⠈⠛⠻⣿⠟⠋⠀⠀⠀⠀⠀⠙⠛⠛⠛⠁⠀⠀⠀⠀⠀⠀⠀⠘⣿⡄⠀⠀⠀⠀
⠀⠀⠀⣰⣿⠷⠶⠖⠒⠲⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠠⠶⠶⠶⠶⣿⣿⣦⠀⠀⠀
⢀⡴⠿⣿⡏⠀⠀⢀⣠⣤⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⡿⠿⢿⣶⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣀⡀⠀⠀⠀⢹⣿⡟⢦⡀
⠉⠀⠸⣿⣇⣤⡾⠛⢉⣠⣶⣤⡀⠀⠀⠀⠀⠀⠀⠀⢠⣴⣿⠟⠁⠀⠀⠀⠈⠛⣷⣦⡀⠀⠀⠀⠀⠀⠀⠀⣠⣤⡀⠉⠛⠷⣦⣄⣸⣿⡇⠀⠈
⠀⠀⠀⣿⣿⡋⠀⠀⣿⣿⡿⠙⠻⢶⣦⣄⠀⠀⠀⠀⠈⠙⣿⣆⠀⠀⠀⠀⠀⣰⡿⠉⠁⠀⠀⠀⢀⣠⣴⡾⠋⢿⣿⣷⠀⠀⠈⢙⣿⣿⡇⠀⠀
⠀⢀⣾⠛⣿⣷⡀⠀⢿⣧⠀⠀⠀⠀⠀⠙⢷⣄⠀⠀⠀⠀⠘⣿⣆⣀⣀⣀⣰⡿⠁⠀⠀⠀⠀⣰⡿⠋⠁⠀⠀⠈⢹⣿⠀⠀⢀⣾⣿⠛⣷⠀⠀
⠀⠸⠁⠀⠈⣿⣿⣶⣼⣿⣇⠀⠀⠀⠀⠀⠈⢿⣆⠀⠀⠀⠀⠀⠙⠛⠋⠛⠋⠀⠀⠀⠀⢀⣾⠟⠀⠀⠀⠀⠀⢀⣿⡏⢀⣤⣾⣟⠁⠀⠈⠇⠀
⠀⠀⠀⠀⠀⣿⡇⠉⠻⣿⡿⠀⠀⠀⠀⠀⠀⠘⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⡏⠀⠀⠀⠀⠀⠀⢸⣿⣶⡿⠛⢻⣿⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢠⣿⡇⠀⣸⣿⠁⠀⠀⠀⠀⠀⠀⠀⢹⣷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⠀⠀⠀⠀⠀⠀⠀⠘⣿⣏⠀⠀⢸⣿⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢸⣿⠀⠐⠿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⠆⠀⠈⣿⡆⠀⠀⠀⠀
⠀⠀⠀⠀⣾⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⣷⠀⠀⠀⠀
⠀⠀⠀⢰⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⡆⠀⠀⠀
⠀⠀⠀⣼⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣧⠀⠀⠀
⠀⠀⢰⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠰⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⠿⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡄⠀⠀
⠀⠀⣼⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡇⠀⠀
⠀⠀⣿⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡇⠀⠀
⠀⠀⢿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⡷⠀⠀
`

const PROMPT = " ⫸⫸ "

var (
	promptStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(""))
	errorStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9"))
	otterStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	outputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	labelStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("13"))
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// env := object.NewEnvironment() // Environment for the REPL session

	for {
		fmt.Fprint(out, promptStyle.Render(PROMPT))
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, otterStyle.Render(OTTER_FACE)+"\n")
	io.WriteString(out, labelStyle.Render(" otter says:")+"\n\n")

	for _, msg := range errors {
		formatted := errorStyle.Render("  ✦ " + msg)
		io.WriteString(out, formatted+"\n")
	}
}
