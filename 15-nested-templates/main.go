package main

import (
	"log/slog"
	"os"
	"text/template"
)

type Score struct {
	One int
	Two int
}

func main() {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		slog.Error("error loading templates", "error", err)
		os.Exit(1)
	}
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", 25); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
