package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"tododly/db"
	"tododly/models"
	"tododly/utils"

	"github.com/golang-jwt/jwt/v5"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := io.ReadAll(r.Body)
	fmt.Println(string(reqBody))

	mapData := make(map[string]string)

	err := json.Unmarshal(reqBody, &mapData)

	if err != nil {
		panic(err)
	}

	email := mapData["email"]
	password := mapData["password"]

	var rowCount int64
	db.Connections.Model(&models.UserCredential{}).Where("email = ?", email).Where("password = ?", password).Count(&rowCount)

	if rowCount == 0 {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("User doesn't exist"))
		panic("User not exists")
	}

	tokenString, err := createToken(email)

	if err != nil {
		panic(err)
	}

	// responseDate := map[string]string{"token": tokenString}

	responseDate, err := json.Marshal(map[string]string{"token": tokenString})
	if err != nil {
		panic(err)
	}

	w.Write(responseDate)
	w.WriteHeader(http.StatusAccepted)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := io.ReadAll(r.Body)
	fmt.Println(string(reqBody))

	userCrendential := models.UserCredential{}
	err := json.Unmarshal([]byte(reqBody), &userCrendential)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(userCrendential)
	}

	// db.Connections.Create(&models.Task())

	// dbTx := db.Connections.Begin()

	// if txnErr != nil {
	// 	panic(txnErr)
	// }

	//Check user existence
	var existingUserCredendials []models.UserCredential

	db.Connections.Where("email = ?", userCrendential.Email).First(&existingUserCredendials)

	fmt.Println("User Exists: ", len(existingUserCredendials) > 0)
	userExists := len(existingUserCredendials) > 0
	// query := `	SELECT COUNT(*)
	// 			FROM user_credentials
	// 			WHERE email = :1    	`

	// // rows, sqlErr := dbTx.Query(query, userCrendential.Email)

	// if sqlErr != nil {
	// 	panic(sqlErr)
	// } else {
	// 	fmt.Println("Data is here")
	// }

	// var count int
	// for rows.Next() {
	// 	if err := rows.Scan(&count); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("count: %d\n", count)
	// }
	// rows.Close()

	// if count == 0 {
	// 	// Create User
	if userExists {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("User exists already."))
		return
	}

	user := createUser(&userCrendential)

	tokenString, err := createToken(user.Email)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic("Token generation successfully")
	}

	user.Token = tokenString

	responseData, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
		return
	}

	fmt.Println("Response Data: ", string(responseData))
	w.Write(responseData)
	w.WriteHeader(http.StatusCreated)

	// 	// if user != nil {
	// 	// 	fmt.Println(user)
	// 	// }

	// 	resBody, _ := json.Marshal(user)
	// 	fmt.Println(string(resBody))
	// 	w.Write(resBody)
	// } else {
	// 	fmt.Println("User already exist")
	// }

	// dbTx.Rollback()
}

func createUser(userCredentials *models.UserCredential) *models.User {
	user := models.User{
		Email:     userCredentials.Email,
		CreatedAt: time.Now(),
		CreatedBy: "new_user",
		UpdatedAt: time.Now(),
		UpdatedBy: "new_user"}

	db.Connections.Create(&user)

	fmt.Println("New User Created: ", user.ID)

	// Create new user
	// insertUserData := `INSERT INTO users(
	// 						user_id,
	// 						email,
	// 						created_at,
	// 						created_by,
	// 						last_updated_at,
	// 						last_updated_by
	// 					) VALUES (
	// 						USER_ID_S.NEXTVAL,
	// 						:1,
	// 						SYSDATE,
	// 						'new_user',
	// 						SYSDATE,
	// 						'new_user'
	// 					) RETURNING
	// 					 	user_id INTO :2`

	// pws, _ := dbTx.Prepare(insertUserData)
	// sqlResult, err := pws.Exec(userCredentials.Email, &userCredentials.UserId)

	// if err != nil {
	// 	fmt.Println(err)
	// 	dbTx.Rollback()
	// 	return nil
	// }

	// rowCount, _ := sqlResult.RowsAffected()
	// fmt.Println("User Inserted = ", rowCount)
	// fmt.Println("User Crendetials :", userCredentials)

	// // Create Auth Data
	// insertAuthData := `INSERT INTO user_credentials(
	// 	user_id,
	// 	email,
	// 	password
	// ) VALUES (
	// 	:1,
	// 	:2,
	// 	:3
	// )`
	userCredentials.UserId = user.ID
	db.Connections.Create(userCredentials)

	// pws, _ = dbTx.Prepare(insertAuthData)
	// sqlResult, err = pws.Exec(userCredentials.UserId, userCredentials.Email, userCredentials.Password)

	// if err != nil {
	// 	fmt.Println(err)
	// 	dbTx.Rollback()
	// 	return nil
	// } else {
	// 	rowCount, _ = sqlResult.RowsAffected()
	// 	fmt.Println("User Credentials Inserted = ", rowCount)
	// }

	// // Get user_data
	// getUserDataQuery := `SELECT * FROM users WHERE user_id = :1`

	// rows, _ := dbTx.Query(getUserDataQuery, userCredentials.UserId)

	// user := models.User{}

	// // if rows.Next() {

	// user = *models.GetUserFromRow(rows)
	// }

	return &user
}

func createToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := claims.SignedString(utils.SecretKey)

	if err != nil {
		return "", err
	}

	fmt.Println("Token generated: ", tokenString)
	return tokenString, nil

}
