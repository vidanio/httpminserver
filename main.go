package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	var err error
	// Logging errors machanism
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fails openning the logging file:", err)
	}
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR :", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Handlers del Servidor HTTP
	s := &http.Server{
		Addr:           ":" + http_port,  // config http port
		Handler:        nil,              // Default Muxer for handler as usual
		ReadTimeout:    20 * time.Second, // send a segment in POST body
		WriteTimeout:   20 * time.Second, // receive a segment in GET req
		MaxHeaderBytes: 1 << 13,          // 8K as Apache and others
	}

	http.HandleFunc("/", root)
	http.HandleFunc(login_cgi, login)
	http.HandleFunc(logout_cgi, logout)

	http.HandleFunc("/contact.cgi", contact)

	log.Fatal(s.ListenAndServe()) // Servidor HTTP multihilo

}
