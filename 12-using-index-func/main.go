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
	var data = map[string]any{
		"words": []string{"zero", "one", "two", "tree", "four", "five"},
		"lname": "Kevin",
	}
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", data); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
