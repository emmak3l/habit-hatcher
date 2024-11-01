package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type habit struct {
	habit     string
	year      int
	month     int
	day       int
	completed bool
}

// AddHabit adds a new habit to the database
func addHabit(db *sql.DB, newHabit habit) {
	// Insert into database, refresh table, etc.
	stmt, _ := db.Prepare("INSERT INTO habit_logs (habit, year, month, day, completed) VALUES (?, ?, ?, ?, ?)")
	stmt.Exec(nil, newHabit.habit, newHabit.year, newHabit.month, newHabit.day, newHabit.completed)
	defer stmt.Close()

	fmt.Printf("Added %v \n", newHabit.habit) // Placeholder function
}

// // MarkHabitComplete marks a habit as complete for a specific day
// func markHabitComplete(db *sql.DB, ourHabit habit) int64 {
// 	// Update database entry for completion, refresh table, etc.
// 	fmt.Println("Marking habit as complete") // Placeholder function
// }

// // UpdateHabitName updates the name of a habit in the database
// func updateHabitName(db *sql.DB, ourHabit habit) int64 {
// 	// Update habit name in database, refresh table, etc.
// 	fmt.Println("Updating habit name") // Placeholder function
// }

// // DeleteHabit deletes a habit from the database
// func deleteHabit(db *sql.DB, habitToDelete string) int64 {
// 	// Delete habit from database, refresh table, etc.
// 	fmt.Println("Deleting a habit") // Placeholder function
// }
