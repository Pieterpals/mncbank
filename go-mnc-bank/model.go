package main

import "time"

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsLogin  bool   `json:"isLogin"`
}

type LoginHistory struct {
	Username string    `json:"username"`
	DateTime time.Time `json:"time"`
}
