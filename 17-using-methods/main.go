package main

import (
	"log/slog"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) SomeProcessing() int {
	return 7
}

func (p Person) AgeDbl() int {
	return p.Age * 2
}

func (p Person) TakesArg(x int) int {
	return x * 2
}

func main() {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		slog.Error("error loading templates", "error", err)
		os.Exit(1)
	}
	var p = Person{
		Name: "Kevin Saucedo",
		Age:  25,
	}
	if err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", p); err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}
}
