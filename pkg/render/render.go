package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(responseWriter http.ResponseWriter, templateName string) {
	templatesDirectory := "./templates/"
	parsedTemplate, err := template.ParseFiles(
		fmt.Sprintf("%s%s", templatesDirectory, templateName),
		fmt.Sprintf("%sbase.layout.gohtml", templatesDirectory))
	if err != nil {
		log.Panic(err)
		return
	}
	err = parsedTemplate.Execute(responseWriter, nil)

	if err != nil {
		log.Println(err)
	}
}
