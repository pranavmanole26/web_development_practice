package templateexample2

import (
	"log"
	"os"
	"text/template"
)

func RenderHtml() {
	tmpl, err := template.ParseFiles("template_example_2/template/one.html")

	if err != nil {
		log.Fatalln("Error parsing the file", err)
	}

	tmpl.Execute(os.Stdout, nil)

	tmpl, err = tmpl.ParseFiles("template_example_2/template/two.html", "template_example_2/template/three.html")

	if err != nil {
		log.Fatalln("Error parsing the file", err)
	}

	log.Println("name", tmpl.Name())

	log.Println("root", tmpl.Root)

	log.Printf("tree: %#v\n", tmpl.Tree)

	if err = tmpl.ExecuteTemplate(os.Stdout, "two.html", nil); err != nil {
		log.Fatalln("Error executing the file", err)
	}
	if err = tmpl.ExecuteTemplate(os.Stdout, "three.html", nil); err != nil {
		log.Fatalln("Error executing the file", err)
	}
}
