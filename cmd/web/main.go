package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
    infoLog *log.Logger
}


func main() {
    addr := flag.String("addr", ":4000", "HTTP network address")
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

    app := application{
        infoLog: infoLog,
    }

    srv := &http.Server{
        Addr: *addr,
        Handler: app.routes(),
        IdleTimeout: time.Minute,
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    err := srv.ListenAndServe()

    if err != nil {
        log.Fatal(err)
    }
}
