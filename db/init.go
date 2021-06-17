package db

import (
	"cohort/global"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var err error

func Init() {
	databaseUrl := "postgres://postgres:postgres@localhost:5438/postgres"
	global.Dbpool, err = pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		panic(err)
	}
	err = global.Dbpool.Ping(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	if err == nil {
		log.Println("Connected to database successfully")
		err = createTables()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("created tables successfully!")
	}
}
