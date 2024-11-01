// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/charmbracelet/bubbles/table"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// 	_ "github.com/mattn/go-sqlite3"
// )

// var baseStyle = lipgloss.NewStyle().
// 	BorderStyle(lipgloss.NormalBorder()).
// 	BorderForeground(lipgloss.Color("240"))

// // initDB function to initialize the SQLite database
// func initDB() (*sql.DB, error) {
// 	// Open or create the SQLite database
// 	db, err := sql.Open("sqlite3", "./habits.db")
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create the habit_logs table if it doesn't already exist
// 	createTableQuery := `
//     CREATE TABLE IF NOT EXISTS habit_logs (
//         id INTEGER PRIMARY KEY AUTOINCREMENT,
//         habit TEXT,
//         year INTEGER,
//         month INTEGER,
//         day INTEGER,
//         completed BOOLEAN,
//         UNIQUE(habit, year, month, day)
//     );
//     `
// 	_, err = db.Exec(createTableQuery)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// func insertSampleData(db *sql.DB) error {
// 	// Sample data for habits and their completed days
// 	habits := []struct {
// 		habit     string
// 		year      int
// 		month     int
// 		day       int
// 		completed bool
// 	}{
// 		{"Exercise", 2024, 11, 1, true},
// 		{"Exercise", 2024, 11, 2, true},
// 		{"Exercise", 2024, 11, 3, false},
// 		{"Meditation", 2024, 11, 1, true},
// 		{"Meditation", 2024, 11, 2, false},
// 		{"Meditation", 2024, 11, 3, true},
// 		{"Reading", 2024, 11, 1, false},
// 		{"Reading", 2024, 11, 2, true},
// 		{"Reading", 2024, 11, 3, true},
// 	}

// 	// Insert sample habits into the database
// 	insertQuery := `
//     INSERT OR IGNORE INTO habit_logs (habit, year, month, day, completed)
//     VALUES (?, ?, ?, ?, ?);
//     `

// 	for _, habit := range habits {
// 		_, err := db.Exec(insertQuery, habit.habit, habit.year, habit.month, habit.day, habit.completed)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // Habit table structure
// func fetchHabitsForMonth(db *sql.DB, year int, month int) ([]table.Row, []table.Column, error) {
// 	query := `
//     SELECT habit, day, completed FROM habit_logs
//     WHERE year = ? AND month = ?
//     ORDER BY habit, day;
//     `
// 	rows, err := db.Query(query, year, month)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	defer rows.Close()

// 	habits := make(map[string][]string)
// 	for rows.Next() {
// 		var habit string
// 		var day int
// 		var completed bool
// 		if err := rows.Scan(&habit, &day, &completed); err != nil {
// 			return nil, nil, err
// 		}

// 		if _, exists := habits[habit]; !exists {
// 			habits[habit] = make([]string, 31) // Create an array of 31 days for each habit
// 		}

// 		habits[habit][day-1] = "✓" // Mark completed days with a tick
// 		if !completed {
// 			habits[habit][day-1] = " " // Blank for incomplete
// 		}
// 	}

// 	// Build table rows
// 	tableRows := []table.Row{}
// 	for habit, days := range habits {
// 		row := table.Row{}

// 		// Append the habit name as the first column (as a string)
// 		row = append(row, habit)

// 		// Append the days of the month (also strings)
// 		row = append(row, days...)

// 		// Add this row to the tableRows slice
// 		tableRows = append(tableRows, row)
// 	}

// 	// Create table columns
// 	tableColumns := []table.Column{
// 		{Title: "Habit", Width: 15},
// 	}
// 	for i := 1; i <= 31; i++ {
// 		tableColumns = append(tableColumns, table.Column{Title: fmt.Sprintf("%d", i), Width: 3})
// 	}

// 	return tableRows, tableColumns, nil
// }

// // Main TUI model
// type model struct {
// 	table table.Model
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	// Handle key messages
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "q", "ctrl+c":
// 			return m, tea.Quit // Exit the program when "q" or "Ctrl+C" is pressed
// 		}
// 	}

// 	// Let the table handle the message
// 	var cmd tea.Cmd
// 	m.table, cmd = m.table.Update(msg)
// 	return m, cmd
// }

// func (m model) View() string {
// 	return baseStyle.Render(m.table.View()) + "\n  " + m.table.HelpView() + "\n"
// }

// func createHabitTable(db *sql.DB) (*table.Model, error) {
// 	currentYear, currentMonth, _ := time.Now().Date()
// 	rows, columns, err := fetchHabitsForMonth(db, currentYear, int(currentMonth))
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create the table model
// 	t := table.New(
// 		table.WithColumns(columns),
// 		table.WithRows(rows),
// 		table.WithFocused(true),
// 		table.WithHeight(7),
// 	)
// 	return &t, nil
// }

// func main() {
// 	// Initialize database
// 	db, err := initDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Insert sample data (optional, for testing)
// 	if err := insertSampleData(db); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create the table view
// 	habitTable, err := createHabitTable(db)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Initialize the Bubble Tea program
// 	p := tea.NewProgram(model{table: *habitTable})
// 	if err := p.Start(); err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/mattn/go-sqlite3"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// initDB function to initialize the SQLite database
func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./habits.db")
	if err != nil {
		return nil, err
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS habit_logs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        habit TEXT,
        year INTEGER,
        month INTEGER,
        day INTEGER,
        completed BOOLEAN,
        UNIQUE(habit, year, month, day)
    );
    `
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func insertSampleData(db *sql.DB) error {
	// Sample data for habits
	habits := []struct {
		habit     string
		year      int
		month     int
		day       int
		completed bool
	}{
		{"Exercise", 2024, 11, 1, true},
		{"Meditation", 2024, 11, 1, true},
		{"Reading", 2024, 11, 1, false},
	}

	insertQuery := `
    INSERT OR IGNORE INTO habit_logs (habit, year, month, day, completed)
    VALUES (?, ?, ?, ?, ?);
    `

	for _, habit := range habits {
		_, err := db.Exec(insertQuery, habit.habit, habit.year, habit.month, habit.day, habit.completed)
		if err != nil {
			return err
		}
	}

	return nil
}

// Habit table structure
func fetchHabitsForMonth(db *sql.DB, year int, month int) ([]table.Row, []table.Column, error) {
	query := `
    SELECT habit, day, completed FROM habit_logs
    WHERE year = ? AND month = ?
    ORDER BY habit, day;
    `
	rows, err := db.Query(query, year, month)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	habits := make(map[string][]string)
	for rows.Next() {
		var habit string
		var day int
		var completed bool
		if err := rows.Scan(&habit, &day, &completed); err != nil {
			return nil, nil, err
		}

		if _, exists := habits[habit]; !exists {
			habits[habit] = make([]string, 31) // Array of 31 days
		}

		if completed {
			habits[habit][day-1] = "✓"
		} else {
			habits[habit][day-1] = " "
		}
	}

	tableRows := []table.Row{}
	for habit, days := range habits {
		row := table.Row{}
		row = append(row, habit)
		row = append(row, days...)
		tableRows = append(tableRows, row)
	}

	tableColumns := []table.Column{
		{Title: "Habit", Width: 15},
	}
	for i := 1; i <= 31; i++ {
		tableColumns = append(tableColumns, table.Column{Title: fmt.Sprintf("%d", i), Width: 3})
	}

	return tableRows, tableColumns, nil
}

// Main TUI model
type model struct {
	table       table.Model
	menu        []string
	menuIdx     int
	db          *sql.DB
	focus       string // "table" or "menu" to keep track of focus
	input       textinput.Model
	addingHabit bool
	completed   bool
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			// Toggle focus between "table" and "menu"
			if m.focus == "table" {
				m.focus = "menu"
			} else {
				m.focus = "table"
			}
		}

		if m.focus == "menu" {
			// Handle menu navigation
			switch msg.String() {
			case "up":
				if m.menuIdx > 0 {
					m.menuIdx--
				}
			case "down":
				if m.menuIdx < len(m.menu)-1 {
					m.menuIdx++
				}
			case "enter":
				// Trigger menu actions based on the selected option
				switch m.menuIdx {
				case 0:
					// addHabit()
				case 1:
					// markHabitComplete()
				case 2:
					// updateHabitName()
				case 3:
					// deleteHabit()
				}
			}
		} else {
			// Let the table handle the navigation
			var cmd tea.Cmd
			m.table, cmd = m.table.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {
	menuStr := "\nMenu (Press Tab to switch focus):\n"
	for i, item := range m.menu {
		cursor := " " // no cursor
		if m.focus == "menu" && m.menuIdx == i {
			cursor = ">" // active cursor for menu focus
		}
		menuStr += fmt.Sprintf("%s %s\n", cursor, item)
	}
	// Add quit instructions at the bottom
	quitInstructions := "\nPress 'q' or 'Ctrl+C' to quit."

	// Show a border or indication for focus if desired
	return baseStyle.Render(m.table.View()) + menuStr + quitInstructions
}

func createHabitTable(db *sql.DB) (*table.Model, error) {
	currentYear, currentMonth, _ := time.Now().Date()
	rows, columns, err := fetchHabitsForMonth(db, currentYear, int(currentMonth))
	if err != nil {
		return nil, err
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	return &t, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := insertSampleData(db); err != nil {
		log.Fatal(err)
	}

	habitTable, err := createHabitTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// Start with the menu in focus
	p := tea.NewProgram(model{
		table: *habitTable,
		menu:  []string{"Add New Habit", "Mark Habit Complete", "Update Habit Name", "Delete Habit"},
		db:    db,
		focus: "menu", // initial focus on the menu
	})
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
