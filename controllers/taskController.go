package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"tododly/db"
	"tododly/models"

	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user_id").(string)
	var tasks []models.Task
	res := db.Connections.Limit(10).Where("user_id = ?", user_id).Find(&tasks)

	if res.Error != nil || len(tasks) == 0 {
		log.Println("Data no found: ", res.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No data available"))
		return
	}

	resBody, err := json.Marshal(&tasks)

	if err != nil {
		log.Println("Error encountered: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resBody)
	w.WriteHeader(http.StatusAccepted)

}

func GetSingleTask(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user_id").(string)
	vars := mux.Vars(r)
	taskID := vars["task_id"]

	var task models.Task
	res := db.Connections.Where("id = ?", taskID).Where("user_id = ?", user_id).First(&task)

	if res.Error != nil {
		log.Println("Error encountered: ", res.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No data available"))
		return
	}

	resBody, err := json.Marshal(&task)

	if err != nil {
		log.Println("Error encountered: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resBody)
	w.WriteHeader(http.StatusAccepted)
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {

	v, _ := io.ReadAll(r.Body)
	username := r.Context().Value("username").(string)

	userCred := models.UserCredential{}
	res := db.Connections.Where("email = ?", username).First(&userCred)

	if res.Error != nil {
		log.Println("Error Encountered: ", res.Error)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unautorized Access"))
		return
	}

	task := models.Task{}
	err := json.Unmarshal([]byte(v), &task)

	if err != nil {
		log.Println("Error Encountered: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request Body"))
		return
	}

	task.CreatedBy = username
	task.UpdatedBy = username
	task.UserId = userCred.UserId

	res = db.Connections.Create(&task)

	if res.Error != nil {
		log.Println("Error creating task: ", res.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	resBody, _ := json.Marshal(&task)
	w.Write(resBody)
	w.WriteHeader(http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	v, _ := io.ReadAll(r.Body)
	taskID := mux.Vars(r)["task_id"]
	username := r.Context().Value("username").(string)

	task := models.Task{}
	err := json.Unmarshal([]byte(v), &task)

	if err != nil {
		log.Println("Error Encountered: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request Body"))
		return
	}

	var mapBody map[string]interface{}
	json.Unmarshal([]byte(v), &mapBody)
	mapBody["updated_by"] = username

	res := db.Connections.Model(&models.Task{}).Where("id = ?", taskID).Updates(mapBody)

	if res.Error != nil {
		log.Println("Error updating task: ", res.Error)
		w.Write([]byte("Error updating task"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res = db.Connections.Where("id = ?", taskID).Find(&task)

	if res.Error != nil {
		log.Println("Error fetching task: ", res.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody, _ := json.Marshal(&task)

	w.Write(resBody)
	w.WriteHeader(http.StatusAccepted)
}
