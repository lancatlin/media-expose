package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

type Session struct {
}

func loadTemplates() {
	tpls = make(map[string]*template.Template)
	mainTmpl := template.New("main")
	mainTmpl.Funcs(template.FuncMap{})
	layoutFiles, err := filepath.Glob("template/layouts/*.html")
	if err != nil {
		panic(err)
	}

	includeFiles, err := filepath.Glob("template/*.html")
	if err != nil {
		panic(err)
	}

	files := make([]string, 1, len(layoutFiles)+1)
	files = append(files, layoutFiles...)
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files[0] = file
		tpl := template.Must(mainTmpl.Clone())
		tpls[strings.TrimSuffix(fileName, ".html")] = template.Must(tpl.ParseFiles(files...))
	}
}

func HTML(w http.ResponseWriter, r *http.Request, page string, data interface{}) {
	pageData := struct {
		Data    interface{}
		Session Session
	}{
		Data:    data,
		Session: Session{},
	}
	if tpl, ok := tpls[page]; ok {
		if err := tpl.ExecuteTemplate(w, "base", pageData); err != nil {
			http.Error(w, err.Error(), 500)
		}
	} else {
		http.Error(w, "cannot find templates", 500)
	}
}

func page(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HTML(w, r, path, nil)
	}
}

func message(w http.ResponseWriter, r *http.Request, title, msg string) {
	HTML(w, r, "message.htm", struct {
		Title, Message string
	}{title, msg})
}
