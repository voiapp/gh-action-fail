package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("terraform", "plan", "--fail")
	_, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌\u00A0 %v\n", err)
		os.Exit(1)
	}
}
