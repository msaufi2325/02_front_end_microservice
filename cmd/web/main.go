package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRootRequest)

	fmt.Println("Starting front end service on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panic(err)
	}
}

func handleRootRequest(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "test.page.gohtml")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	// Define the paths to the partial templates
	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	// Create a slice to hold the template paths
	templatePaths := []string{
		fmt.Sprintf("./cmd/web/templates/%s", templateName),
	}

	// Append the partial templates to the slice
	templatePaths = append(templatePaths, partials...)

	// Parse the template files
	tmpl, err := template.ParseFiles(templatePaths...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with no data
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
