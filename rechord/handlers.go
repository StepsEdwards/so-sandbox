package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := templateData {
		Page: "Home",
	}

	app.render(w, "home.tmpl.html", &data)
}

func (app *application) account(w http.ResponseWriter, r *http.Request) {
	data := templateData {
		Page: "Account",
	}

	app.render(w, "account.tmpl.html", &data)
}

func (app *application) progressionCreate(w http.ResponseWriter, r *http.Request) {
	data := templateData {
		Page: "Progression Create",
	}

	app.render(w, "chord-progression.tmpl.html", &data)
}

func (app *application) progressionCreatePost(w http.ResponseWriter, r *http.Request) {
	data := templateData {
		Page: "Progression Post",
	}

	app.render(w, "chord-progression.tmpl.html", &data)
}

func (app *application) education(w http.ResponseWriter, r *http.Request) {
	data := templateData {
		Page: "Education",
	}

	app.render(w, "education.tmpl.html", &data)
}

func (app *application) educationGlossary(w http.ResponseWriter, r *http.Request) {
	data := templateData {
		Page: "Glossary",
	}

	app.render(w, "glossary.tmpl.html", &data)
}

func (app *application) educationScales(w http.ResponseWriter, r *http.Request) {
	data := templateData {
		Page: "Scales",
	}

	app.render(w, "scales.tmpl.html", &data)
}

func (app *application) educationChords(w http.ResponseWriter, r *http.Request) {
	data := templateData {
		Page: "Chords",
	}

	app.render(w, "chords.tmpl.html", &data)
}

func (app *application) educationCircleOf5ths(w http.ResponseWriter, r *http.Request) {
	data := templateData {
		Page: "Circle of 5ths",
	}

	app.render(w, "circle-of-5ths.tmpl.html", &data)
}
