package models

import (
	"database/sql"
)

/*
CREATE TABLE `users` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`password` varchar(256) NOT NULL,
	`login_id` varchar(32) NOT NULL,
	`name` varchar(32) DEFAULT NULL,
	`email` varchar(128) DEFAULT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY `email` (`email`)
  ) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/

type User struct {
	ID       int            `gorm:"primaryKey;column:id;" json:"id"`
	LoginID  string         `gorm:"column:login_id" json:"login_id"`
	Password string         `gorm:"column:password" json:"password"`
	Name     sql.NullString `gorm:"column:name"     json:"name"`
	Email    sql.NullString `gorm:"column:email"    json:"email"`
}

func (User) TableName() string {
	return "users"
}

type AddUserData struct {
	LoginID  string `json:"login_id" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	LoginID  string `json:"login_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResult struct {
	UserID       int    `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
