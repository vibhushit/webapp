package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/vibhushit/webapp/pkg/config"
	"github.com/vibhushit/webapp/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// method a bit complex but magical
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	//if i am in development environment I dont want to use cache and build template every time
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	// parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	// err := parsedTemplate.Execute(w, nil)

	// if err != nil {
	// 	fmt.Println("Error parsing template", err)
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil

}

//simpler method
// caching the templates
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {

// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tc[t]

// 	if !inMap {
// 		//is template is not in the map create the template
// 		log.Println("creating template")
// 		err = createTemplateCache(t)

// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {

// 		//we have the template in the cache
// 		log.Println("using cached template")

// 	}
// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl
// 	return nil
// }
