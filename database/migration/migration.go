package migration

import (
	"database/sql"

	"github.com/ariopri/Let-It-Be/tree/main/backend/database/seeder"
	_ "github.com/go-sql-driver/mysql"
)

func Migrate(db *sql.DB) {
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id INTEGER PRIMARY KEY AUTO_INCREMENT,
				firstname VARCHAR(255) NOT NULL,
				lastname VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				password VARCHAR(255) NOT NULL,
				role VARCHAR(255) CHECK (role IN ('admin', 'user')) DEFAULT 'user',
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);`)
	if err != nil {
		panic(err)
	}
	seeder.Seed(db)
}
