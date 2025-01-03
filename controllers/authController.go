package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
	"tododly/db"
	"tododly/models"
	"tododly/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := io.ReadAll(r.Body)
	log.Println(string(reqBody))

	mapData := make(map[string]string)

	err := json.Unmarshal(reqBody, &mapData)

	if err != nil {
		http.Error(w, "Unable to read request", http.StatusBadRequest)
		log.Println("Unable to read request: ", err)
		return
	}

	email := mapData["email"]
	password := mapData["password"]

	userCred := models.UserCredential{}

	res := db.Connections.Where("email = ?", email).First(&userCred)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("User doesn't exist"))
		log.Println("User not exists")
		return
	}

	isValidPassword, err := verifyPassword(password, userCred.Password)

	if !isValidPassword || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect username/password"))
		log.Println("Incorrect password")
		return
	}

	tokenString, err := createToken(userCred.Email, userCred.UserId)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Token generation failed: %v", err)
		return
	}

	responseDate, err := json.Marshal(map[string]string{"token": tokenString})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Response body encoding failed: %v", err)
		return
	}

	w.Write(responseDate)
	w.WriteHeader(http.StatusAccepted)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := io.ReadAll(r.Body)
	log.Println(string(reqBody))

	userCrendential := models.UserCredential{}
	err := json.Unmarshal([]byte(reqBody), &userCrendential)

	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Println("Unable to read request: ", err)
		return
	}

	//Check user existence
	var existingUserCredendials []models.UserCredential

	db.Connections.Where("email = ?", userCrendential.Email).First(&existingUserCredendials)

	log.Println("User Exists: ", len(existingUserCredendials) > 0)
	userExists := len(existingUserCredendials) > 0

	// Create User
	if userExists {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("User exists already."))
		log.Println("User exists already.")
		return
	}

	user, err := createUser(&userCrendential)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Response body encoding failed: %v", err)
		return
	}

	log.Println("Response Data: ", string(responseData))
	w.Write(responseData)
	w.WriteHeader(http.StatusCreated)

}

func createUser(userCredentials *models.UserCredential) (*models.User, error) {

	hashedPassword, err := generatePasswordHash(userCredentials.Password)

	if err != nil {
		return nil, err
	}

	userCredentials.Password = string(hashedPassword)

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

	log.Println("New User Created: ", user.ID)

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

	tokenString, err := claims.SignedString(utils.JWT_SECRET_KEY)

	if err != nil {
		return "", err
	}

	log.Println("Token generated: ", tokenString)
	return tokenString, nil

}

func generatePasswordHash(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Printf("Password verification failed: %v", err)
		return "", err
	}

	return string(hash), nil
}

func verifyPassword(password string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		log.Printf("Password verification error: %v", err)
		return false, err
	}
	return true, nil
}
