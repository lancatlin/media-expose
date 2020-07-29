package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

type Meta struct {
}

func (h *Handler) loadTemplates() {
	h.tpls = make(map[string]*template.Template)
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
		h.tpls[strings.TrimSuffix(fileName, ".html")] = template.Must(tpl.ParseFiles(files...))
	}
}

func (h Handler) HTML(w http.ResponseWriter, r *http.Request, page string, data interface{}) {
	pageData := struct {
		Data interface{}
		Meta Meta
	}{
		Data: data,
		Meta: Meta{},
	}
	if tpl, ok := h.tpls[page]; ok {
		if err := tpl.ExecuteTemplate(w, "base", pageData); err != nil {
			http.Error(w, err.Error(), 500)
		}
	} else {
		http.Error(w, "cannot find templates", 500)
	}
}

func (h Handler) page(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.HTML(w, r, path, nil)
	}
}

func (h Handler) message(w http.ResponseWriter, r *http.Request, title, msg string) {
	h.HTML(w, r, "message.htm", struct {
		Title, Message string
	}{title, msg})
}
