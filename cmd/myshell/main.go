package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func handleCdCommand(args []string) {
	path := args[0]
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", path)
		return
	}

	if !fileInfo.IsDir() {
		fmt.Fprintf(os.Stdout, "cd: %s: Not a directory\n", path)
		return
	}

	os.Chdir(path)
}

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
			statusCode = 128
		}
	}
	os.Exit(statusCode)
}

func handleTypeCommand(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stdout, "type: missing argument")
		return
	}

	envPath := os.Getenv("PATH")
	envPaths := strings.Split(envPath, ":")

	shellBuiltinCommands := []string{"cd", "echo", "exit", "type"}
	searchCommand := args[0]
	for _, shellBuiltinCommand := range shellBuiltinCommands {
		if searchCommand == shellBuiltinCommand {
			fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", searchCommand)
			return
		}

		for _, path := range envPaths {
			if _, err := os.Stat(path + "/" + searchCommand); err == nil {
				fmt.Fprintf(os.Stdout, "%s is %s/%s\n", searchCommand, path, searchCommand)
				return
			}
		}
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", searchCommand)
}

func invokeShellCommand(command string, args ...string) (output string, err error) {
	envPath := os.Getenv("PATH")
	envPaths := strings.Split(envPath, ":")

	for _, path := range envPaths {
		if _, err := os.Stat(path + "/" + command); err == nil {
			cmd := exec.Command(command, args...)
			var out strings.Builder
			cmd.Stdout = &out
			if err := cmd.Run(); err != nil {
				return "", err
			}

			return out.String(), nil
		}
	}

	return "", fmt.Errorf("%s: not found", command)
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		rawInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stdout, "failed to parse input: %s\n", err)
			break
		}

		trimmedInput := strings.TrimSpace(rawInput)

		input := strings.Split(trimmedInput, " ")
		command := input[0]
		args := input[1:]

		switch command {
		case "cd":
			handleCdCommand(args)
		case "echo":
			handleEchoCommand(args)
		case "exit":
			handleExitCommand(args)
		case "type":
			handleTypeCommand(args)
		default:
			output, err := invokeShellCommand(command, args...)
			if err != nil {
				fmt.Fprintln(os.Stdout, err)
				break
			}

			fmt.Fprint(os.Stdout, output)
		}
	}
}
