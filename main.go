package main

import (
	"html/template" // Para renderizar arquivos HTML como templates
	"log"

	// Para exibir logs no terminal
	"net/http" // Para gerenciar as requisições HTTP
	// Para converter strings em números
	// Para manipulação de strings (usado na busca)
	"sync" // Para sincronizar o acesso a dados compartilhados (uso de mutex)
)

type Task struct {
	ID        int
	Name      string
	Completed bool
}

var (
	tasks    = []Task{}
	nextID   = 1
	mu       sync.Mutex
	tpml     = template.Must(template.ParseGlob("templates/*.html"))
	pageSize = 5
)

func listTasks(w http.ResponseWriter, r *http.Request) {
	// A função está vazia por enquanto
}

func createTask(w http.ResponseWriter, r *http.Request) {
	// A função está vazia por enquanto
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	// A função está vazia por enquanto
}

func editTask(w http.ResponseWriter, r *http.Request) {
	// A função está vazia por enquanto
}

func completeTask(w http.ResponseWriter, r *http.Request) {
	// A função está vazia por enquanto
}

func searchTasks(w http.ResponseWriter, r *http.Request) {
	// A função está vazia por enquanto
}

func main() {
	http.HandleFunc("/", listTasks)
	http.HandleFunc("/create", createTask)
	http.HandleFunc("/edit", editTask)
	http.HandleFunc("/delete", deleteTask)
	http.HandleFunc("/complete", completeTask)
	http.HandleFunc("/search", searchTasks)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Servidor rodando na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
