package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	ConfPath  string
	configure Configure
)

func main() {
	flag.StringVar(&ConfPath, "conf", "configure.yaml", "-conf to specify configure file.")
	var err error
	configure, err = LoadConfigure(ConfPath)
	if err != nil {
		InitialConfigure(ConfPath)
		return
	}
	fmt.Println(configure)
	handler := NewHandler(OpenDB(configure))
	srv := http.Server{
		Handler:      handler.router,
		Addr:         fmt.Sprintf("0.0.0.0:%d", configure.Server.Port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	fmt.Printf("Server listen on %s\n", configure.Server.Base)
	log.Fatal(srv.ListenAndServe())
}
