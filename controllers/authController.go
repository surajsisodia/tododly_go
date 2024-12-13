package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"tododly/db"
	"tododly/models"
	"tododly/utils"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
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

	userCred := models.UserCredential{}

	res := db.Connections.Where("email = ?", email).Where("password = ?", password).First(&userCred)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("User doesn't exist"))
		panic("User not exists")
	}

	tokenString, err := createToken(userCred.Email, userCred.UserId)

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

	user, err := createUser(&userCrendential)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseData, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	fmt.Println("Response Data: ", string(responseData))
	w.Write(responseData)
	w.WriteHeader(http.StatusCreated)

}

func createUser(userCredentials *models.UserCredential) (*models.User, error) {

	tokenString, err := createToken(userCredentials.Email, userCredentials.UserId)

	if err != nil {
		return nil, errors.New("token generation failed")
	}

	user := models.User{
		Email:     userCredentials.Email,
		Token:     &tokenString,
		CreatedAt: time.Now(),
		CreatedBy: "new_user",
		UpdatedAt: time.Now(),
		UpdatedBy: "new_user"}

	db.Connections.Create(&user)

	fmt.Println("New User Created: ", user.ID)

	userCredentials.UserId = user.ID
	db.Connections.Create(&userCredentials)

	return &user, nil
}

func createToken(username string, userId int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"user_id":  userId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := claims.SignedString(utils.SecretKey)

	if err != nil {
		return "", err
	}

	fmt.Println("Token generated: ", tokenString)
	return tokenString, nil

}
