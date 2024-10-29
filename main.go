package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

// initDB function to initialize the SQLite database
func initDB() (*sql.DB, error) {
	// Open or create the SQLite database
	db, err := sql.Open("sqlite3", "./habits.db")
	if err != nil {
		return nil, err
	}

	// Create the habit_logs table if it doesn't already exist
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
	// Sample data for habits and their completed days
	habits := []struct {
		habit     string
		year      int
		month     int
		day       int
		completed bool
	}{
		{"Exercise", 2024, 10, 1, true},
		{"Exercise", 2024, 10, 2, true},
		{"Exercise", 2024, 10, 3, false},
		{"Meditation", 2024, 10, 1, true},
		{"Meditation", 2024, 10, 2, false},
		{"Meditation", 2024, 10, 3, true},
		{"Reading", 2024, 10, 1, false},
		{"Reading", 2024, 10, 2, true},
		{"Reading", 2024, 10, 3, true},
	}

	// Insert sample habits into the database
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
			habits[habit] = make([]string, 31) // Create an array of 31 days for each habit
		}

		habits[habit][day-1] = "✓" // Mark completed days with a tick
		if !completed {
			habits[habit][day-1] = " " // Blank for incomplete
		}
	}

	// Build table rows
	tableRows := []table.Row{}
	for habit, days := range habits {
		row := table.Row{}

		// Append the habit name as the first column (as a string)
		row = append(row, habit)

		// Append the days of the month (also strings)
		row = append(row, days...)

		// Add this row to the tableRows slice
		tableRows = append(tableRows, row)
	}

	// Create table columns
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
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Handle key messages
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit // Exit the program when "q" or "Ctrl+C" is pressed
		}
	}

	// Let the table handle the message
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.table.View()
}

func createHabitTable(db *sql.DB) (*table.Model, error) {
	currentYear, currentMonth, _ := time.Now().Date()
	rows, columns, err := fetchHabitsForMonth(db, currentYear, int(currentMonth))
	if err != nil {
		return nil, err
	}

	// Create the table model
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	return &t, nil
}

func main() {
	// Initialize database
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	// Insert sample data (optional, for testing)
	if err := insertSampleData(db); err != nil {
		log.Fatal(err)
	}

	// Create the table view
	habitTable, err := createHabitTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the Bubble Tea program
	p := tea.NewProgram(model{table: *habitTable})
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
