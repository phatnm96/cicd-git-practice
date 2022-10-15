package main

import (
	"log"
	"os"

	"cutloss-trading/app/db"
	"cutloss-trading/app/routers"
)

func main() {
	db, err := db.InitStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}

	echo := routers.Init(db)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	echo.Logger.Fatal(echo.Start(":" + httpPort))
}
