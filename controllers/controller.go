package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tododly/db"
	"tododly/models"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {

	tasks := []models.Task{}

	db.Connections.Find(&tasks)

	for _, user := range tasks {
		fmt.Println("Title", user.Title)
	}

	resBody, _ := json.Marshal(&tasks)

	w.Write(resBody)
	w.WriteHeader(http.StatusAccepted)

	// w.Write([]byte("Getting all tasks"))

	// query := `SELECT * FROM tasks`

	// // pws, _ := db.Connections.Prepare(query)

	// // // sqlResult, err := pws.Exec()

	// // if err != nil {
	// // 	fmt.Println("Error in sql result :", err)
	// // }

	// sqlResult, _ := db.Connections.Query(query)
	// // var name string

	// var (
	// 	task_id         int
	// 	title           string
	// 	description     string
	// 	created_at      time.Time
	// 	created_by      string
	// 	last_updated_at time.Time
	// 	last_updated_by string
	// )

	// for sqlResult.Next() {
	// 	sqlResult.Scan(&task_id, &title, &description, &created_at, &created_by, &last_updated_at, &last_updated_by)

	// 	fmt.Println("Row Data: ", task_id, " ", title, " ", description)
	// }

}

func GetSingleTask(w http.ResponseWriter, r *http.Request) {

}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("Insiding Create New Task")

	// v, _ := io.ReadAll(r.Body)
	// fmt.Println(string(v))

	// task := models.Task{}
	// err := json.Unmarshal([]byte(v), &task)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Decoded JSON :", task)

	// insertQuery := `INSERT INTO tasks
	//                (TASK_ID,
	//                 TITLE,
	//                 DESCRIPTION,
	//                 CREATED_AT,
	//                 CREATED_BY,
	//                 LAST_UPDATED_AT,
	//                 LAST_UPDATED_BY)

	//             VALUES(
	//                 TASK_ID_S.NEXTVAL,
	//                 :1,
	//                 :2,
	//                 SYSDATE,
	//                 'user',
	//                 SYSDATE,
	//                 'user'
	//             )`

	// pws, _ := db.Connections.Prepare(insertQuery)
	// sqlResult, err := pws.Exec(task.Title, task.Description)

	// if err != nil {
	// 	fmt.Println("Error in sql result", err)
	// }

	// rowCount, _ := sqlResult.RowsAffected()
	// fmt.Println("Inserted number of rows = ", rowCount)
	// w.Write([]byte("Insert Task successfull"))
}
