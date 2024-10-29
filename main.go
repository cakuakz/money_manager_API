package main

import (
	"log"
	"money-manager/config"
	"os"

	"money-manager/migrate"
	"money-manager/routers"
)

func main() {
	db := config.SetUpDatabaseConnection()
    defer config.CloseDatabaseConnection(db)

    if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
            if arg == "--migrate" {
                if err := migrate.Migrate(db); err != nil {
                    log.Fatalf("error migration: %v", err)
                }
                log.Println("migration completed successfully")
            }
        }
	}

    var PORT = ":8080"

    routers.StartServer(db).Run(PORT)
}