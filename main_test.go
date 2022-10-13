package main

import (
	"fmt"
	"testing"
)

func TestReadArgs(t *testing.T) {
	params := []string{"30", "echo", `"hola`, `mundo"`}
	wanta, wantb, wantc := 30, "echo", `"hola mundo"`
	a, b, c := ReadArgs(params)
	if a != wanta || b != wantb || c != wantc {
		t.Fatalf(`ReadArgs(30,"echo", "hola mundo")=%v,%v,%v want match for %v, %v, %v`, a, b, c, wanta, wantb, wantc)
	}
}

func TestExecuteCommand(t *testing.T) {
	pa := "echo"
	pb := "hola mundo"
	wa := "hola mundo\n"
	a := ExecuteCommand(pa, pb)
	fmt.Println(a)
	if a != wa {
		t.Fatalf(`ExecuteCommand("echo hola mundo")=%v, want match for %v`, a, wa)
	}

}
