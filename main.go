package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path is required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// arrow keys for histroy, tab for autocomplete

func main() {
	reader := bufio.NewReader(os.Stdin)
	histroy := make([]string, 0)

	for {
		pwd, _ := exec.Command("pwd").Output()
		pwd = pwd[:len(pwd)-1]

		fmt.Printf("--> %s ", pwd)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if input == "" {
			continue
		}

		if err = execCommand(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		histroy = append(histroy, input)
	}
}
