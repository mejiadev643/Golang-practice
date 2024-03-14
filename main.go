package main

import (
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/mejiadev/go-gorm/DB"
	"github.com/mejiadev/go-gorm/models"
	"github.com/mejiadev/go-gorm/routes"
)

func main() {
	db.Connect()

	db.DB.AutoMigrate(models.User{}, models.Task{})
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//tasks
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTasksHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":8085", r)
}
