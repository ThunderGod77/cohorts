package db

import (
	"cohort/global"
	"context"
	"log"
)

func createTables() error {
	_, err := global.Dbpool.Exec(context.Background(), "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err != nil {
		return err
	}
	commandTag, err := global.Dbpool.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users ("+
		"id uuid DEFAULT uuid_generate_v4 (),"+
		"username VARCHAR NOT NULL UNIQUE,"+
		"email VARCHAR NOT NULL UNIQUE,"+
		"password VARCHAR NOT NULL,"+
		"payment BOOLEAN NOT NULL,"+
		"join_date DATE NOT NULL DEFAULT CURRENT_DATE,"+
		"auth VARCHAR NOT NULL DEFAULT 'student',"+
		"payment_date DATE,"+
		"PRIMARY KEY (id)"+
		");")
	log.Println(commandTag)
	if err != nil {
		return err
	}
	return nil
}
