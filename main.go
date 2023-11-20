package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

var todos []Todo

func todoHandler(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title: "TODO List",
		Todos: todos,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func addHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	newTask := r.FormValue("newTask")
	if newTask != "" {
		todos = append(todos, Todo{Item: newTask, Done: false})
	}

	http.Redirect(w, r, "/todo", http.StatusSeeOther)

}

func updateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	index, err := strconv.Atoi(r.URL.Path[len("/update/"):])

	if err != nil || index < 0 || index >= len(todos) {
		http.Error(w, "Invalid task index", http.StatusBadRequest)
		return
	}
	updatedTask := r.FormValue("updateTask")
	done := r.FormValue("done") == "on"
	todos[index].Item = updatedTask
	todos[index].Done = done
	http.Redirect(w, r, "/todo", http.StatusSeeOther)

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(r.URL.Path[len("/delete/"):])

	if err != nil || index < 0 || index >= len(todos) {
		http.Error(w, "Invalid task index", http.StatusBadRequest)
		return
	}

	todos = append(todos[:index], todos[index+1:]...)

	http.Redirect(w, r, "/todo", http.StatusSeeOther)

}
func main() {
	todos = []Todo{
		{Item: "Install GO", Done: true},
		{Item: "Learn Go", Done: false},
	}
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todoHandler)
	mux.HandleFunc("/add", addHandler)
	mux.HandleFunc("/update/", updateHandler)
	mux.HandleFunc("/delete/", deleteHandler)
	http.ListenAndServe(":8080", mux)
}
