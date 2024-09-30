package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n  " + m.table.HelpView() + "\n"
}

func main() {
	columns := []table.Column{
		{Title: "Habit", Width: 15},
		{Title: "1", Width: 2},
		{Title: "2", Width: 2},
		{Title: "3", Width: 2},
		{Title: "4", Width: 2},
		{Title: "5", Width: 2},
		{Title: "6", Width: 2},
		{Title: "7", Width: 2},
		{Title: "8", Width: 2},
		{Title: "9", Width: 2},
		{Title: "10", Width: 2},
		{Title: "11", Width: 2},
		{Title: "12", Width: 2},
		{Title: "13", Width: 2},
		{Title: "14", Width: 2},
		{Title: "15", Width: 2},
		{Title: "16", Width: 2},
		{Title: "17", Width: 2},
		{Title: "18", Width: 2},
		{Title: "19", Width: 2},
		{Title: "20", Width: 2},
		{Title: "21", Width: 2},
		{Title: "22", Width: 2},
		{Title: "23", Width: 2},
		{Title: "24", Width: 2},
		{Title: "25", Width: 2},
		{Title: "26", Width: 2},
		{Title: "27", Width: 2},
		{Title: "28", Width: 2},
		{Title: "29", Width: 2},
		{Title: "30", Width: 2},
		{Title: "31", Width: 2},
	}

	rows := []table.Row{
		{"Brush Teeth", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X"},
		{"Floss Teeth", "X", "", "X", "", "X", "X", "X", "", "X", "X", "X", "", "X", "", "X", "X", "X", "", "X", "X", "X", "", "X", "", "X", "X", "X", "", "X", "X", ""},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
