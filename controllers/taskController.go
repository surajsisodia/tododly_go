package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tododly/db"
	"tododly/models"

	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []models.Task
	res := db.Connections.Limit(10).Find(&tasks)

	if res.Error != nil || len(tasks) == 0 {
		fmt.Println("Data no found: ", res.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No data available"))
	}

	resBody, _ := json.Marshal(&tasks)

	w.Write(resBody)
	w.WriteHeader(http.StatusAccepted)

}

func GetSingleTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one item")

	vars := mux.Vars(r)
	taskID := vars["task_id"]

	var task models.Task
	db.Connections.Where("id = ?", taskID).First(&task)

	fmt.Println("Task Found: ", task)

	if task.Title == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resBody, err := json.Marshal(&task)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.Write(resBody)
	w.WriteHeader(http.StatusAccepted)
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {

	v, _ := io.ReadAll(r.Body)
	fmt.Println(string(v))

	username := r.Context().Value("username").(string)

	task := models.Task{}
	err := json.Unmarshal([]byte(v), &task)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	task.CreatedBy = username
	task.UpdatedBy = username

	db.Connections.Create(&task)

	// fmt.Println("Task Created with task_id: ", task.Id)

	resBody, _ := json.Marshal(&task)
	w.Write(resBody)
	w.WriteHeader(http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	v, _ := io.ReadAll(r.Body)
	fmt.Println(string(v))
	taskID := mux.Vars(r)["task_id"]
	username := r.Context().Value("username").(string)

	task := models.Task{}
	err := json.Unmarshal([]byte(v), &task)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var mapBody map[string]interface{}
	json.Unmarshal([]byte(v), &mapBody)
	mapBody["updated_by"] = username

	res := db.Connections.Model(&models.Task{}).Where("id = ?", taskID).Updates(mapBody)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res = db.Connections.Where("id = ?", taskID).Find(&task)

	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody, _ := json.Marshal(&task)

	w.Write(resBody)
	w.WriteHeader(http.StatusAccepted)
}
