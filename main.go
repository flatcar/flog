package main

import (
	"github.com/flatcar-linux/flog/app"
	"github.com/flatcar-linux/flog/pkg/db"
)

func main() {
	// create the main app
	app := &app.App{}

	// create the DB repository
	db := &db.MockDB{}

	app.Initialize(db)
	app.Run(":3000")
}
