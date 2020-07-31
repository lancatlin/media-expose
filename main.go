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

	ConfPath            string
	ImportCompaniesPath string
	ImportMediaPath     string
	configure           Configure
)

func main() {
	parseFlags()
	var err error
	configure, err = LoadConfigure(ConfPath)
	if err != nil {
		InitialConfigure(ConfPath)
		fmt.Printf("configure file has been successfully generated at %s\n", ConfPath)
		return
	}

	OpenDB(configure)

	if ImportCompaniesPath != "" {
		ImportCSVFromFile(ImportCompaniesPath, ImportCompanies)
	}
	if ImportMediaPath != "" {
		ImportCSVFromFile(ImportMediaPath, ImportMedia)
	}
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

func parseFlags() {
	flag.StringVar(&ConfPath, "conf", "configure.yaml", "-conf to specify configure file.")
	flag.StringVar(&ImportCompaniesPath, "import-companies", "", "-import-companies to specify the companies CSV file path.")
	flag.StringVar(&ImportMediaPath, "import-media", "", "-import-media to specify the media CSV file path.")
	flag.Parse()
}
