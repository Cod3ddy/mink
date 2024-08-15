package main

import (
	"crypto/tls"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/form/v4"
)

type application struct {
	Logger      *slog.Logger
	TemplateCache map[string]*template.Template
	FormDecoder *form.Decoder
}

func main() {
	address := flag.String("address", ":3000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	templateCache, err := newTemplateChache()
	if err != nil{
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	app := application{
		Logger: logger,
		TemplateCache: templateCache,
		FormDecoder: formDecoder,
	}
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	 server := &http.Server{
		Addr: *address,
		Handler: app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig: tlsConfig,
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server on %s", "addr", server.Addr)
	err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")

	logger.Error(err.Error())
	os.Exit(1)
}