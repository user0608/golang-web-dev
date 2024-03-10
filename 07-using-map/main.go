package main

import (
	"log/slog"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		slog.Error("error loading templates", "error", err)
		os.Exit(1)
	}
	// ejecuta el primero que encuentra
	// var sages = []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}
	var sages = map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammd",
	}
	if err := tpl.Execute(os.Stdout, sages); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
