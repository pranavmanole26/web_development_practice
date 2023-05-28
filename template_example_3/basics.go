package templateexample3

import (
	"log"
	"os"
	"text/template"
)

func RenderHtml() {
	tmpl, err := template.ParseGlob("template_example_3/template/*.html")
	if err != nil {
		log.Fatalln("Error parsing html", err)
	}

	if err = tmpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln("Error executing template", err)
	}

	if err = tmpl.ExecuteTemplate(os.Stdout, "one.html", nil); err != nil {
		log.Fatalln("Error executing template", err)
	}

	if err = tmpl.ExecuteTemplate(os.Stdout, "two.html", nil); err != nil {
		log.Fatalln("Error executing template", err)
	}

	if err = tmpl.ExecuteTemplate(os.Stdout, "three.html", nil); err != nil {
		log.Fatalln("Error executing template", err)
	}
}
