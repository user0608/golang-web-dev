package main

import (
	"log/slog"
	"os"
	"text/template"
)

type User struct {
	Name  string
	Motto string
	Admin bool
}

func (u User) String() string {
	return u.Name
}

func main() {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		slog.Error("error loading templates", "error", err)
		os.Exit(1)
	}
	var users = []User{
		{
			Name:  "Buddha",
			Motto: "The belief of no beliefs",
		},
		{
			Name:  "Gandhi",
			Motto: "Be the change",
			Admin: true,
		},
		{
			Name:  "",
			Motto: "Nobody",
			Admin: true,
		},
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", users); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
