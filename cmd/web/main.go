package main

import (
	"flag"
	"log"
	"net/http"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	errorLog := log.New(log.Writer(), "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(log.Writer(), "INFO\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: log.New(log.Writer(), "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Handler:  app.routes(),
	}

	log.Printf("Starting server on %s", *addr)

	err := srv.ListenAndServe()
	log.Fatal(err)
}
