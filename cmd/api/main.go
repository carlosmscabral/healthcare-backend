package main

import (
	"log"
	"net/http"
	"os"
)

const defaultAddr = ":8080"

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	addr := defaultAddr
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Print("Starting server on ", addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
