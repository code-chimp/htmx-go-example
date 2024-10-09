package main

import (
	"github.com/code-chimp/htmx-go-example/internal/models"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type templateData struct {
	CurrentYear int
	Contact     *models.Contact
	Contacts    []*models.Contact
	Form        any
}

// humanDate returns a human readable string representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// functions is a map of functions that can be used in templates.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// newTemplateCache creates a new template cache.
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	err := filepath.Walk("./ui/html/pages", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".gohtml") {
			relPath, err := filepath.Rel("./ui/html/pages", path)
			if err != nil {
				return err
			}
			key := strings.ReplaceAll(relPath, string(filepath.Separator), ".")
			name := filepath.Base(path)

			// parse base layout to create a new template set
			ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/layouts/base.gohtml")
			if err != nil {
				return err
			}

			// parse all partials and add them to the template set
			// ts, err = ts.ParseGlob("./ui/html/partials/*.gohtml")
			// if err != nil {
			// 	return err
			// }

			// finally parse the page template and add it to the template set
			ts, err = ts.ParseFiles(path)
			if err != nil {
				return err
			}

			cache[key] = ts
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return cache, nil
}
