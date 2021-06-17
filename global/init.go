package global

import (
	"log"
	"os"
)

func Init() {
	var err error
	File, err = os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(File, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(File, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(File, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
