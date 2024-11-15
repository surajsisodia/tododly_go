package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//UserId       uint           `gorm:"primaryKey;autoIncrement" json:"user_id"`
	ID           int            `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName    string         `json:"first_name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	PhotoUrl     string         `json:"photo_url,omitempty"`
	Email        string         `json:"email"`
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    string         `json:"created_by"`
	UpdatedAt    time.Time      `json:"last_updated_at"`
	UpdatedBy    string         `json:"last_updated_by"`
	DeletedAt    gorm.DeletedAt `gorm:"-"`
}

// func GetUserFromRow(rows *sql.Rows) *User {
// 	user := User{}
// 	// var badDataError bool = false

// 	mapData := utils.ConvertRowsToMap(rows)

// 	fmt.Println(mapData)
// 	fmt.Println("Testing : ", mapData[0]["user_id"])

// 	// userMapData := mapData[0]

// 	// user_id, er := strconv.Atoi(userMapData["user_id"])
// 	// if er != nil {
// 	// 	return nil, "BAD_DATA"
// 	// }

// 	// if _ , er := strconv.Atoi(userMapData["user_id"]), err != nil {

// 	// }

// 	fmt.Printf("%T: %s\n", mapData[0]["user_id"], mapData[0]["user_id"])

// 	// user := User{UserId: strconv.Atoi(userMapData["user_id"])}

// 	// columns, _ := row.Columns()
// 	// fmt.Println("Columns:")
// 	// for _, colName := range columns {
// 	// 	fmt.Println(colName)
// 	// }
// 	// byteData, _ := json.Marshal(mapData[0])
// 	// err := json.Unmarshal(byteData, &user)
// 	// err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.PhotoUrl, &user.Token, &user.RefreshToken, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// fmt.Println(user)
// 	return &user
// }
