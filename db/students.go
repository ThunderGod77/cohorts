package db

import (
	"cohort/global"
	"context"
	"github.com/jackc/pgx/v4"
)

//to verify if there is an existing student or not

type studentData struct {
	ID    string
	Passw string
	Auth  string
}

func GetStudent(email string, username string) (*studentData, bool, error) {
	var id string
	var passw string
	var auth string
	//query to get students with given email or username
	err := global.Dbpool.QueryRow(context.Background(), "SELECT id,password,auth FROM users WHERE username=$1 OR email=$2", username, email).Scan(&id, &passw, &auth)
	if err == pgx.ErrNoRows {
		return nil, false, nil
	} else {
		sD := studentData{
			ID:    id,
			Passw: passw,
			Auth:  auth,
		}
		return &sD, true, err
	}

}

// to add student to the database
func AddStudent(email, username, password string) error {
	//query to insert students
	_, err := global.Dbpool.Exec(context.Background(), "INSERT INTO users (username,email,password,payment) VALUES ($1,$2,$3,$4)", username, email, password, false)
	return err
}
