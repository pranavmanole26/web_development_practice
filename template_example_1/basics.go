package templateexample1

import (
	"log"
	"os"
	"text/template"
)

func RenderHtml() {
	tmpl, err := template.ParseFiles("template_example_1/template/one.html")

	if err != nil {
		log.Fatalln("Error parsing the file", err)
	}

	tmpl.Execute(os.Stdout, nil)
}
