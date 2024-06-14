package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleEchoCommand(args []string) {
	fmt.Fprintln(os.Stdout, strings.Join(args, " "))
}

func handleExitCommand(args []string) {
	var err error
	statusCode := 0
	if len(args) > 0 {
		statusCode, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stdout, "failed to parse exit status code:", err)
			statusCode = 1
		}
	}
	os.Exit(statusCode)
}

func handleTypeCommand(args []string) {
	shellBuiltinCommands := []string{"echo", "exit", "type"}
	for _, shellBuiltinCommand := range shellBuiltinCommands {
		if len(args) > 0 && args[0] == shellBuiltinCommand {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", args[0])
			return
		}
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", args[0])
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
		command := input[0]
		args := input[1:]

		switch command {
		case "echo":
			handleEchoCommand(args)
		case "exit":
			handleExitCommand(args)
		case "type":
			handleTypeCommand(args)
		default:
			output := fmt.Sprintf("%s: command not found\n", command)
			fmt.Fprint(os.Stdout, output)
		}
	}
}
