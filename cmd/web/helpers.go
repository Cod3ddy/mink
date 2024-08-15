package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error){
	var (
		method = r.Method
		uri = r.URL.RequestURI()
	)

	app.Logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}


func (app *application) clientError(w http.ResponseWriter, status int){
	http.Error(w, http.StatusText(status), status)
}


// Render templates [ui]

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData){
	ts, ok := app.TemplateCache[page]
	if !ok{
		err := fmt.Errorf("template %s does not exist", page)
		app.serverError(w,r,err)
		
	}

	// init new buffer

	buffer := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buffer, "base", data)
	if err != nil{
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	buffer.WriteTo(w)
}

func (app *application) newTemplateData(r *http.Request) templateData{
	return templateData{}
}
