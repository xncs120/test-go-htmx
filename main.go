package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type CmdNote struct {
	Command string
	Explain string
}

func list(w http.ResponseWriter, r *http.Request) {
	cmdNotes := map[string][]CmdNote{
		"cmdNotes": {
			{Command: "go run main.go", Explain: "run go app"},
			{Command: "go mod init dir_path", Explain: "generate go.mod"},
			{Command: "go mod tidy", Explain: "generate go.sum"},
		},
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, cmdNotes)
}

func saveList(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(1 * time.Second)
	command := r.PostFormValue("command")
	explain := r.PostFormValue("explain")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "command-row", CmdNote{Command: command, Explain: explain})
}

func main() {
	http.HandleFunc("/", list)
	http.HandleFunc("/add-command", saveList)
	port := ":8000"
	fmt.Printf("Server running on port %s ...", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
