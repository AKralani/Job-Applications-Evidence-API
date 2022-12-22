package main

import (
	"log"
	"net/http"

	"github.com/akralani/jobapplications/database"
	"github.com/akralani/jobapplications/evidence"
	_ "github.com/go-sql-driver/mysql"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	evidence.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
