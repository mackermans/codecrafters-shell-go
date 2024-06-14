package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("error: ", err)
	}

	output := fmt.Sprintf("%s: command not found\n", strings.TrimSpace(input))
	fmt.Fprint(os.Stdout, output)
}
