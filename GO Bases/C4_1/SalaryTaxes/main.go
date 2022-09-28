package main

import (
	"bufio"
	"fmt"
	"os"
)

type taxError struct {
	message   string
	codeError int
}

func (err *taxError) NewTaxError() string {

}

func main() {
	reader := bufio.NewReader(os.Stdout)
	fmt.Print("> ")
	salary, _ := reader.ReadString('\n')

}
