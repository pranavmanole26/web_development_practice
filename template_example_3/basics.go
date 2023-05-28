package templateexample3

import (
	"log"
	"os"
	"text/template"
)

func RenderHtml() {
	//Parses all the templates which matches the pattern and stores into the template container
	tmpl, err := template.ParseGlob("template_example_3/template/*.html")
	if err != nil {
		log.Fatalln("Error parsing html", err)
	}

	//Executes first template found in the template container
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
