package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("failed to parse input: ", err)
			break
		}

		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "exit" {
			break
		}

		output := fmt.Sprintf("%s: command not found\n", trimmedInput)
		fmt.Fprint(os.Stdout, output)
	}
}
