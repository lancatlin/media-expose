package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB

	tpls map[string]*template.Template

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
	OpenDB(configure)
	loadTemplates()
	r := NewRouter()
	srv := http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%d", configure.Server.Port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	fmt.Printf("Server listen on %s\n", configure.Server.Base)
	log.Fatal(srv.ListenAndServe())
}
