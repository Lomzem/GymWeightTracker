package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type MovementData struct {
	Movement     string
	MaxWeight    int
	TargetMuscle string
}

func main() {
	exampleMvData := map[string][]MovementData{
		"MovementData": {
			{Movement: "Bench Press", MaxWeight: 135, TargetMuscle: "Chest"},
			{Movement: "Back Squad", MaxWeight: 225, TargetMuscle: "Quads"},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

		tmpl.Execute(w, exampleMvData)
		// tmpl.ExecuteTemplate()
	})

	http.HandleFunc("/add-movement", func(w http.ResponseWriter, r *http.Request) {
		mvmtName := r.PostFormValue("movement-name")

		// TODO: handle non-int weight
		mvmtWeight, _ := strconv.Atoi(r.PostFormValue("movement-weight"))
		mvmtTarget := r.PostFormValue("movement-target")

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "movement", MovementData{
			Movement:     mvmtName,
			MaxWeight:    mvmtWeight,
			TargetMuscle: mvmtTarget,
		})
	})

	http.HandleFunc("/remove-movement", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("YOU RELALY WANT THIS?!?")
	})

	log.Fatal(http.ListenAndServe(":6969", nil))
}
