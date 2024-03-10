package main

import (
	"log/slog"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		slog.Error("error loading templates", "error", err)
		os.Exit(1)
	}
	// ejecuta el primero que encuentra
	if err := tpl.Execute(os.Stdout, nil); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "tpl2.html", nil); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
