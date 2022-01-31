package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Company  string `json:"company"`
}

func main() {
	user1 := &User{
		ID:       1,
		Name:     "Md Abu Raihan",
		Username: "raihan",
		Email:    "raihanmahmudi35@gmail.com",
		Address:  "kushtia",
		Phone:    "01853566901",
		Website:  "mdabraihan.tech",
		Company:  "Master Academy",
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
