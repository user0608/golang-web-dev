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
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatterd ...",
	}
	var sages = []sage{buddha, gandhi, mlk}
	if err := tpl.Execute(os.Stdout, sages); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
