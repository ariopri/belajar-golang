package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ariopri/Let-It-Be/tree/main/backend/database/migration"
	"github.com/ariopri/Let-It-Be/tree/main/backend/handler"
	"github.com/ariopri/Let-It-Be/tree/main/backend/handler/middleware"
	"github.com/ariopri/Let-It-Be/tree/main/backend/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	//database
	db, err := sql.Open("mysql", "root:tanahdamai@tcp(localhost:3306)/pusing")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migration.Migrate(db)

	r := gin.Default()

	//middleware
	m := middleware.InitMiddleware()
	r.Use(m.CORS())

	// users
	u := repository.NewUserRepo(db)
	handler.NewUserHandler(r, u)

	r.Run()
}
