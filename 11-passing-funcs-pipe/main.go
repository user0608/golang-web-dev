package main

import (
	"log/slog"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

func Split(s string) []string {
	return strings.Split(s, "")
}

func Join(ss []string) string {
	return strings.Join(ss, "-")
}

var fn = template.FuncMap{
	"uc": strings.ToUpper,
	"sp": Split,
	"jn": Join,
}

func main() {
	tpl, err := template.New("").Funcs(fn).ParseFiles("tpl.html")
	if err != nil {
		slog.Error("error loading templates", "error", err)
		os.Exit(1)
	}
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", buddha); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
