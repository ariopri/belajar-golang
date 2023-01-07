package seeder

import (
	"database/sql"
	"log"

	"github.com/ariopri/Let-It-Be/tree/main/backend/utils/hash"
)

func Seed(db *sql.DB) {

	hashedPassword, err := hash.HashPassword("Password@123")

	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
		INSERT INTO users (firstname, lastname, email, password, role)
		VALUES (?, ?, ?, ?, ?)
	`, "John", "Doe", "john@doe.com", hashedPassword, "user")
	if err != nil {
		log.Fatal(err)
	}
}
