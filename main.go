package main

import (
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
	ticker := time.NewTicker(1 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			if <-done {
			} else {
				_ = <-ticker.C
				ExecuteCommand(command, args)
			}
		}
	}()
}

func ReadArgs(args []string) (int, string, string) {
	second, _ := strconv.Atoi(args[0])
	program := args[1]
	programArgs := strings.Join(args[2:], " ")
	return second, program, programArgs
}

// main Programa principal
func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("timer <seconds> <command> <params>")
		os.Exit(1)
	}
	seconds, program, programArgs := ReadArgs(args[1:])
	run(seconds, program, programArgs)

}
