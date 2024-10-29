package migrate

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sql_migration
var dbMigrations embed.FS

var DbConnection *sql.DB

func Migrate(db *sql.DB) error {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migration",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	DbConnection = db

	fmt.Println("Migration success, applied", n, "migrations!")
	return nil
}
