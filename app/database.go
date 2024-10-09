package app

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_api_boilerplate")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// migrate -database "mysql://root@tcp(localhost:3306)/golang_api_boilerplate" -path db/migrations up
	// migrate -database "mysql://root@tcp(localhost:3306)/golang_api_boilerplate" -path db/migrations down
	// migrate -database "mysql://root@tcp(localhost:3306)/golang_api_boilerplate" -path db/migrations down 2

	//Dirty State (stuck tidak bisa up dan down)
	// migrate -database "mysql://root@tcp(localhost:3306)/golang_api_boilerplate" -path db/migrations version
	// migrate -database "mysql://root@tcp(localhost:3306)/golang_api_boilerplate" -path db/migrations force *version*
}
