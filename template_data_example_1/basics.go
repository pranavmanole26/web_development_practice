package templatedataexample1

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("template_data_example_1/*.html"))
}

func RenderHtml() {
	if err := tmpl.ExecuteTemplate(os.Stdout, "index.html", "from another world"); err != nil {
		log.Fatalln("Error executing template", err)
	}
}
