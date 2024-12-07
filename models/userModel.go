package models

import (
	"time"
)

type User struct {
	ID           int       `gorm:"primaryKey" json:"user_id"`
	FirstName    *string   `json:"first_name,omitempty"`
	LastName     *string   `json:"last_name,omitempty"`
	PhotoUrl     *string   `json:"photo_url,omitempty"`
	Email        string    `gorm:"not null" json:"email"`
	Token        string    `gorm:"not null" json:"token"`
	RefreshToken string    `gorm:"not null" json:"refresh_token"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	CreatedBy    string    `gorm:"not null" json:"created_by"`
	UpdatedAt    time.Time `gorm:"not null" json:"last_updated_at"`
	UpdatedBy    string    `gorm:"not null" json:"last_updated_by"`
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
