package main

import (
	"os/exec"
)

func main() {
	exec.Command("terraform", "plan", "--fail")
}