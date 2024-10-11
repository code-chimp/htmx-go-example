package main

import (
	"flag"
	"fmt"
	"github.com/code-chimp/htmx-go-example/internal/services"
	"github.com/code-chimp/htmx-go-example/internal/vcs"
	"github.com/go-playground/form/v4"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

var revision = vcs.Revision()

// application struct holds the application-wide dependencies.
type application struct {
	logger      *slog.Logger
	contacts    *services.ContactRepository
	templates   map[string]*template.Template
	formDecoder *form.Decoder
}

func main() {
	addr := flag.Int("addr", 4000, "HTTP network address")
	displayVersion := flag.Bool("version", false, "Display version information")
	flag.Parse()

	if *displayVersion {
		fmt.Printf("HTMX Demo Website:\n\tVersion:\t%s\n\tRevison:\t%s\n", version, revision)
		os.Exit(0)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	contactRepository, err := services.NewRepository()
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
		Addr:         fmt.Sprintf(":%d", *addr),
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
			fmt.Sprintf("http://localhost%s", srv.Addr),
		),
	)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
