package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tododly/db"
	"tododly/models"
)

func GetMyProfile(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user_id").(string)

	user := models.User{}

	res := db.Connections.Where("id = ?", user_id).First(&user)

	if res.Error != nil {
		fmt.Println(res.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error Encountered"))
		return
	}

	resBody, _ := json.Marshal(&user)

	w.WriteHeader(http.StatusAccepted)
	w.Write(resBody)
}

func UpdateMyProfile(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	v, _ := io.ReadAll(r.Body)
	fmt.Println(string(v))

	task := models.Task{}
	err := json.Unmarshal([]byte(v), &task)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var mapBody map[string]interface{}
	json.Unmarshal([]byte(v), &mapBody)
	mapBody["updated_by"] = r.Context().Value("username").(string)

	res := db.Connections.Model(&models.User{}).Where("id = ?", user_id).Updates(mapBody)

	if res.Error != nil {
		fmt.Println(res.Error)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	var user models.User
	res = db.Connections.Where("id = ?", user_id).First(&user)

	if res.Error != nil {
		fmt.Println(res.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resBody, _ := json.Marshal(&user)
	w.WriteHeader(http.StatusAccepted)
	w.Write(resBody)
}
