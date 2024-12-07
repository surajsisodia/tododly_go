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

	//Check user existence
	var existingUserCredendials []models.UserCredential

	db.Connections.Where("email = ?", userCrendential.Email).First(&existingUserCredendials)

	fmt.Println("User Exists: ", len(existingUserCredendials) > 0)
	userExists := len(existingUserCredendials) > 0

	// Create User
	if userExists {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("User exists already."))
		return
	}

	user := createUser(&userCrendential)

	tokenString, err := createToken(user.Email)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic("Token generation failed")
	}

	user.Token = tokenString

	responseData, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	fmt.Println("Response Data: ", string(responseData))
	w.Write(responseData)
	w.WriteHeader(http.StatusCreated)

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

	userCredentials.UserId = user.ID
	db.Connections.Create(userCredentials)

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
