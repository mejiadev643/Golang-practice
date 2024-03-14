package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/mejiadev/go-gorm/DB"
	"github.com/mejiadev/go-gorm/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("Task not found"))
		return

	}
	json.NewEncoder(w).Encode(&task)
}
func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Task
	json.NewDecoder(r.Body).Decode(&tasks)
	createTask := db.DB.Create(&tasks)
	if createTask.Error != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(createTask.Error.Error()))
		return
	}
	json.NewEncoder(w).Encode(&tasks)
}
func UpdateTasksHandler(w http.ResponseWriter, r *http.Request) {
}
func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("Task not found"))
		return
	}
	db.DB.Delete(&task) //eliminado logico
	//json.NewEncoder(w).Encode("Task deleted successfully")
	w.WriteHeader(http.StatusNoContent) //204 eliminado no content
}
