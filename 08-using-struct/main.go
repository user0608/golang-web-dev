package main

import (
	"log/slog"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

func main() {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		slog.Error("error loading templates", "error", err)
		os.Exit(1)
	}
	// ejecuta el primero que encuentra
	// var sages = []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	if err := tpl.Execute(os.Stdout, buddha); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
