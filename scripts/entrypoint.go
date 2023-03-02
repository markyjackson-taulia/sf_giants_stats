package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter the team's abbreviation
	fmt.Print("Enter the team's abbreviation (e.g. SF for San Francisco Giants): ")
	teamAbbr, _ := reader.ReadString('\n')
	teamAbbr = strings.TrimSpace(teamAbbr)

	// Call the myapp executable with the user's input
	cmd := exec.Command("/app/myapp", teamAbbr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running myapp: %v\n", err)
		os.Exit(1)
	}
}
