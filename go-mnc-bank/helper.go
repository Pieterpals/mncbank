package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func BoolIsRegisteredUser(username string) bool {

	jsonFile, err := os.Open("jsonFile/user.json")

	if err != nil {
		return true
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users
	var flag bool

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Username == username {
			flag = true
			break
		} else {
			flag = false
		}
	}

	return flag
}

func GetPassword(username string) string {
	jsonFile, err := os.Open("jsonFile/user.json")

	if err != nil {
		return "failed open file JSON"
	}

	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users Users
	var password string
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Username == username {
			password = users.Users[i].Password
		} else {
			password = ""
		}
	}

	return password
}

func GetUser(username string) User {
	jsonFile, err := os.Open("jsonFile/user.json")

	if err != nil {
		return User{}
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	req := User{}

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Username == username {
			req.ID = users.Users[i].ID
			req.Username = users.Users[i].Username
			req.Password = users.Users[i].Password
			req.IsLogin = users.Users[i].IsLogin
		}
	}

	return req
}
