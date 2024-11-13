package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"groupie-tracker/backend"
)

func main() {
	// Parse command-line flag for address
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	// Setup leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Setup HTTP server with routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", backend.ArtistsHandler)
	mux.HandleFunc("/artist/", backend.ArtistHandler)
	mux.HandleFunc("/location/", backend.LocationHandler)
	mux.HandleFunc("/date/", backend.DateHandler)
	mux.HandleFunc("/relation/", backend.RelationHandler)
	mux.HandleFunc("/about", backend.AboutHandler)

	serv := http.Server{
		Addr:     *addr,
		Handler:  mux,
		ErrorLog: errorLog,
	}

	// Start the server in a goroutine in a separate routine
	go func() {
		infoLog.Printf("Starting server on http://localhost%s", *addr)
		err := serv.ListenAndServe()
		if err != nil {
			errorLog.Fatal(err)
		}
	}()

	// Block main routine to prevent program exit
	select {} // Keeps the main function alive indefinitely without consuming cpu cycles
}
