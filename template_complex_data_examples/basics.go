package templatecomplexdataexamples

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type Person struct {
	Name string
	Age  int
}

func init() {
	tmpl = template.Must(template.ParseGlob("template_complex_data_examples/templates/*.html"))
}

func RenderHtml() {
	sl := []string{"Pranav", "Aishwarya", "?"}
	if err := tmpl.ExecuteTemplate(os.Stdout, "slice.html", sl); err != nil {
		log.Fatalln("Error executing template", err)
	}

	mp := map[string]string{
		"name1": "Pranav",
		"name2": "Aishwarya",
		"name3": "?",
	}
	if err := tmpl.ExecuteTemplate(os.Stdout, "map.html", mp); err != nil {
		log.Fatalln("Error executing template", err)
	}

	strct := Person{
		Name: "Pranav",
		Age:  30,
	}
	if err := tmpl.ExecuteTemplate(os.Stdout, "struct.html", strct); err != nil {
		log.Fatalln("Error executing template", err)
	}
}
