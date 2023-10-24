package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}

// caching the templates
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {

	var tmpl *template.Template
	var err error

	_, inMap := tc[t]

	if !inMap {
		//is template is not in the map create the template
		log.Println("creating template")
		err = createTemplateCache(t)

		if err != nil {
			log.Println(err)
		}
	} else {

		//we have the template in the cache
		log.Println("using cached template")

	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}
