package main

import (
	"net/http"

	"github.com/cod3ddy/mink/internal/validator"
)

type shortenUrlForm struct{
	URL string `form:"url"`
	validator.Validator `form:"-"`
}


func (app *application) home(w http.ResponseWriter, r *http.Request){
	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) shortenUrl(w http.ResponseWriter, r *http.Request){
	var form shortenUrlForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.URL), "url", "url cannot be blank!")

	if !form.Valid(){
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w,r, http.StatusUnprocessableEntity, "home.html", data)
		return
	}

	w.Write([]byte("shortening url....."))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}