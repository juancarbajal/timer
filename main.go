package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func ExecuteCommand(command string, args string) string {
	cmd := exec.Command(command, args)

	stdout, err := cmd.Output()
	if err != nil {
		return err.Error()
	} else {
		return string(stdout)
	}
}

// run ...
func run(seconds int, command string, args string) {
	for {
		fmt.Println(ExecuteCommand(command, args))
		time.Sleep(time.Duration(seconds) * time.Second)
	}
}

func ReadArgs(args []string) (int, string, string, error) {
	second, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, "", "", errors.New("second must be integer")
	}
	program := args[1]
	if len(args) <= 2 {
		return second, program, "", nil
	}
	programArgs := strings.Join(args[2:], " ")
	return second, program, programArgs, nil
}

// main Programa principal
func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("timer <seconds> <command> <params>")
		os.Exit(1)
	}
	seconds, program, programArgs, err := ReadArgs(args[1:])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	run(seconds, program, programArgs)
}
