package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	state "github.com/mrabasto/user-logger/internal/app"
)

func initialModel() state.Model {
	return state.Model{
		Choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		Selected: make(map[int]struct{}),
	}
}

func main() {
	app := tea.NewProgram(initialModel())
	if _, err := app.Run(); err != nil {
		fmt.Printf("Alas, there's an error: %v", err)
		os.Exit(1)
	}
}
