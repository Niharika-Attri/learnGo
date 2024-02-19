package main

import (
	"github.com/Niharika-Attri/learnGo/16-packages/db"
	errHandler "github.com/Niharika-Attri/learnGo/16-packages/errorHandler"
)

func main() {
	errHandler.PrintError("Something went wrong", 401)

	db.ConnectToDB()
}
