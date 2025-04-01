package state

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		case "enter", " ":
			_, ok := m.Selected[m.Cursor]

			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}

		}
	}

	return m, nil
}

func (m Model) View() string {
	s := "What should we buy in the market?\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)

	}

	s += "\nPress q to quit.\n"

	return s
}
