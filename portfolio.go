//package main

import (
	"html/template"
	"log"
	"net/http"
)

type Project struct {
	Name        string
	Description string
}

type Portfolio struct {
	Name     string
	Projects []Project
}

func main() {
	portfolio := Portfolio{
		Name: "My Portfolio",
		Projects: []Project{
			{Name: "Project 1", Description: "This is the first project."},
			{Name: "Project 2", Description: "This is the second project."},
			{Name: "Project 3", Description: "This is the third project."},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index").Parse(`
			<!DOCTYPE html>
			<html>
				<head>
					<title>{{.Name}}</title>
				</head>
				<body>
					<h1>{{.Name}}</h1>
					<ul>
						{{range .Projects}}
							<li><strong>{{.Name}}</strong>: {{.Description}}</li>
						{{end}}
					</ul>
				</body>
			</html>
		`)
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, portfolio)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
