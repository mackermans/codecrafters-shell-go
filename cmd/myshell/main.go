package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleExitCommand(args []string) {
	var err error
	statusCode := 0
	if len(args) > 0 {
		statusCode, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("failed to parse exit status code:", err)
			statusCode = 1
		}
	}
	os.Exit(statusCode)
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("failed to parse input:", err)
			break
		}

		trimmedInput := strings.TrimSpace(rawInput)

		input := strings.Split(trimmedInput, " ")

		switch command := input[0]; command {
		case "exit":
			handleExitCommand(input[1:])
		default:
			output := fmt.Sprintf("%s: command not found\n", command)
			fmt.Fprint(os.Stdout, output)
		}
	}
}
