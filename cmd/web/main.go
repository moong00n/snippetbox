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

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: log.New(log.Writer(), "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Handler:  mux,
	}

	log.Printf("Starting server on %s", *addr)

	err := srv.ListenAndServe()
	log.Fatal(err)
}
