package backend

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Set the model for the choices
type model struct {
	choices  []string         // items in the Todo list
	cursor   int              // which todo list item is under the cursor
	selected map[int]struct{} // which todo list item is selected
}

// Initialice the model with data
func InitalModle() model {
	return model{
		choices: []string{"golang", "express", "java spring"},
		// IDK
		selected: make(map[int]struct{}),
	}
}

// Init the model
func (m model) Init() tea.Cmd {
	return nil
}

// Update the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// cursor starts at 0
		// also starts at the top so to move you need to ++ the int
		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			fmt.Println(m.choices[m.cursor])
		fmt.Println("selected")

		}
	}
	return m, nil
}

// what gets redert in the terminal
func (m model) View() string {
	s := "Select a Backend Framework\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s \n", cursor, choice)
	}
	s += "\nPress q to quit. \n"
	return s
}
