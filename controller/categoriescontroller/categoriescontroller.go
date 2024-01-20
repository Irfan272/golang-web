package categoriescontroller

import (
	"html/template"
	"net/http"
	"time"
	"web-crud/entities"
	"web-crud/model/categorymodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/categories/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/categories/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var categories entities.Categories

		// Parse the form values
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		categories.Name = r.FormValue("name")
		categories.CreatedAt = time.Now()
		categories.UpdatedAt = time.Now()

		if ok := categorymodel.Create(categories); !ok {
			temp, _ := template.ParseFiles("views/categories/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}

}
func Edit(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
