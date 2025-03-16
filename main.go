package main

import (
	"fmt"
	"os"
	"strings"

	dbManager "github.com/audible-sound/rpr-go-server/db"
	"github.com/audible-sound/rpr-go-server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	db := dbManager.SetupDatabase()

	if len(os.Args) > 1 {
		var command string = strings.ToLower(os.Args[1])
		if command == "migrate" {
			dbManager.MigrateTables(db)
		} else if command == "drop" {
			dbManager.DropTables(db)
		} else {
			fmt.Println("Error: Command does not exist")
		}
		return
	}

	// Start the server
	router := gin.Default()
	router.Run(":3000")
}
