package main

import (
	"flag"
	"fmt"
	"github.com/code-chimp/htmx-go-example/internal/models"
	"github.com/go-playground/form/v4"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// application struct holds the application-wide dependencies.
type application struct {
	logger      *slog.Logger
	contacts    *models.ContactRepository
	templates   map[string]*template.Template
	formDecoder *form.Decoder
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	contactRepository, err := models.NewRepository()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	app := &application{
		logger:      logger,
		contacts:    contactRepository,
		templates:   templateCache,
		formDecoder: formDecoder,
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info(
		"starting server",
		slog.String(
			"addr",
			fmt.Sprintf("http://localhost%s", *addr),
		),
	)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
