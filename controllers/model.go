package controllers

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserType  int    `json:"usertype"`
}

type UserResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    []User `json:"data"`
}
