package main

import (
	"exchange-rate/intrenal/app"
	"exchange-rate/intrenal/di"
	"exchange-rate/intrenal/storage"
	"log"
)

func main() {
	container := di.BuildContainer()

	var appInstance *app.App
	err := container.Invoke(func(app *app.App) {
		appInstance = app
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Run the application
	if err := appInstance.Run(); err != nil {
		log.Fatalf("Error running app: %v", err)
	} else {
		log.Println("App ran successfully")
	}

	// Close the database connection
	err = container.Invoke(func(storage *storage.Storage) {
		if closeErr := storage.Db.Close(); closeErr != nil {
			log.Printf("Error closing database: %v", closeErr)
		}
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
