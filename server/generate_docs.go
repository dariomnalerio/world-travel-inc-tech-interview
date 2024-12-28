package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("swag", "init", "-g", "./cmd/api/main.go", "-o", "./docs")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error generating docs: %v", err)
	}
}
