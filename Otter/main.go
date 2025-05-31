package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/charmbracelet/lipgloss"
	"github.com/r-priyanshu/interpreter/repl"
)

const otterBanner = `
        ....              s         s                            
    .x~X88888Hx.         :8        :8                            
   H8X 888888888h.      .88       .88                  .u    .   
  8888:"*888888888:    :888ooo   :888ooo      .u     .d88B :@8c  
  88888:        "%8  -*8888888 -*8888888   ud8888.  ="8888f8888r 
. "88888          ?>   8888      8888    :888'8888.   4888>'88"  
". ?888%           X   8888      8888    d888 '88%"   4888> '    
  ~*??.            >   8888      8888    8888.+"      4888>      
 .x88888h.        <   .8888Lu=  .8888Lu= 8888L       .d888L .+   
:"""8888888x..  .x    ^%888*    ^%888*   '8888c. .+  ^"8888*"    
"    "*888888888"       'Y"       'Y"     "88888%       "Y"      
        ""***""                             "YP'                 
                                                                 
                                                                 
`

func main() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user: %v", err)
	}


	bannerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("63")). 
		Bold(true).
		MarginTop(1)

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).
		Bold(true)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("42"))

	dividerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		SetString("--------------------------------------------------")

	fmt.Println(bannerStyle.Render(" " + otterBanner))

	fmt.Println(labelStyle.Render(" Logged in as: ") + valueStyle.Render(currentUser.Username))
	fmt.Println(dividerStyle)

	repl.Start(os.Stdin, os.Stdout)
}
