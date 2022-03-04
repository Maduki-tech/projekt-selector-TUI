package main

import (
	"fmt"
	"os"

	"david/project-selector/frontend"
	tea "github.com/charmbracelet/bubbletea"
)
func main() {
	p := tea.NewProgram(frontend.InitalModle())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
