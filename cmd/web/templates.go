package main

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/cod3ddy/mink/ui"
)


type templateData struct{
	Form any
	Status string
}

func newTemplateChache() (map[string]*template.Template, error){
	// Init cahe map
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil{
		return nil, err 
	}


	for _, page := range pages{
		name := filepath.Base(page)

		patterns  := []string{
			"html/base.html",
			page, 
		}

		ts, err := template.New(name).ParseFS(ui.Files, patterns...)

		if err != nil{
			return nil, err
		}
		
		cache [name] = ts
	}

	return cache, nil
}