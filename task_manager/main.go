package main

import (
	"ayushbhargav/task_manager/cmd"
	"ayushbhargav/task_manager/db"
)

func main() {
	db.Init()
	cmd.Execute()
	// db.CloseDB()
}
