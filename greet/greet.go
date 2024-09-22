package greet

import (
	"bufio"
	"fmt"
	"io"
)
func GreetUser(input io.Reader, output io.Writer){
    name := "you"
    fmt.Fprintln(output, "What is your name?")
    userInput := bufio.NewScanner(input)
    if userInput.Scan(){
        name = userInput.Text()
    }
    fmt.Fprintf(output, "Hello, %s.\n", name)
}