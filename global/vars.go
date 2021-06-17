package global

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	File          *os.File
)

var Dbpool *pgxpool.Pool