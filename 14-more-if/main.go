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
	var score = Score{
		One: 7,
		Two: 9,
	}
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", score); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
