package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type application struct {
	templateCache map[string]*template.Template
}

type templateData struct {
	Chords []string
	Page string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"./templates/base.tmpl.html",
			page,
		}

		tmpl, err := template.New(name).ParseFiles(patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}

func (app *application) render(w http.ResponseWriter, page string, data *templateData) {
	tmpl, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("The template %s does not exist", page)
		panic(err)
	}

	buffer := new(bytes.Buffer)

	if err := tmpl.ExecuteTemplate(buffer, "base", data); err != nil {
		panic(err.Error())
	}

	buffer.WriteTo(w)
}

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("GET /account", app.account)
	mux.HandleFunc("GET /progression/create", app.progressionCreate)
	mux.HandleFunc("POST /progression/create", app.progressionCreatePost)
	mux.HandleFunc("GET /education", app.education)
	mux.HandleFunc("GET /education/glossary", app.educationGlossary)
	mux.HandleFunc("GET /education/scales", app.educationScales)
	mux.HandleFunc("GET /education/chords", app.educationChords)
	mux.HandleFunc("GET /education/circle-of-5ths", app.educationCircleOf5ths)

	return mux
}

func main()  {
	addr := flag.String("addr", ":4000", "Http network address")


	templateCache, err := newTemplateCache()
	if err != nil {
		panic(err.Error())
	}

	app := &application{
		templateCache: templateCache,
	}
	
	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(),
	}

	fmt.Printf("Starting server on %s\n", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
