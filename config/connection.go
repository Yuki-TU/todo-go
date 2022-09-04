package config

import (
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/Yuki-TU/todo-go/utils"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

// DBへの初回接続
//
// @reuturn db データベースインスタンス
func Connection() *sql.DB {
	databaseURI := make(chan string, 1)
	databaseURI <- utils.GodotEnv("DATABASE_URI")

	db, err := sql.Open("mysql", <-databaseURI)

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	// マイグレーション
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	m.Up()

	return db
}
