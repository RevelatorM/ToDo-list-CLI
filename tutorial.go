package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// =========================
func addTask(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin) //creating reader Stdin - standart input,bufio - library (package) gives NewReader

	fmt.Print("add task: ")
	task, _ := reader.ReadString('\n')

	stmt, err := db.Prepare("INSERT INTO tasks(task, status) VALUES(?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}

	stmt.Exec(task, "todo")
	fmt.Println("task was added!")
}

// =========================
func showTask(db *sql.DB) {
	rows, err := db.Query("SELECT id, task, status FROM tasks") //db.Query returns rows
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close() //defer calls rows.Close() at the end of the code

	fmt.Println("\n task list:")
	for rows.Next() {
		var id int
		var task, status string

		rows.Scan(&id, &task, &status)
		fmt.Printf("%d | %s | %s\n", id, task, status) // %d — number (id), %s — string (task and status)
	}
}

// =========================
func deleteTask(db *sql.DB) {
	fmt.Print("write task`s ID: ")
	var id int
	fmt.Scan(&id)

	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		fmt.Println(err)
		return
	}

	stmt.Exec(id)
	fmt.Println("the task was deleted")
}

//=========================

func main() {
	database, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	//=========================
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS tasks(id INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT, status TEXT)")
	if err != nil {
		fmt.Println("error while preparation (Prepare):", err)
		return
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Println("error while making a request (Exec):", err)
		return
	}
	statement.Close()
	//=========================
	fmt.Println("Welcome to Todo list \n")
	//=========================
	for {
		fmt.Println("\n1 - add task")
		fmt.Println("2 - show tasks")
		fmt.Println("3 - delete tasks")
		fmt.Println("0 - exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTask(database)
		case 2:
			showTask(database)
		case 3:
			deleteTask(database)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("error")
		}
	}
	//=========================
}
