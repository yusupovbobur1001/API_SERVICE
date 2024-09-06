package models

type Response struct {
	StatusCode  int         `json:"status_code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

type User struct {
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

type UpdateTask struct {
	UserId string `json:"user_id"`
	Title string `json:"title"`
}

type UpdateDetail struct {
	TaskId string `json:"task_id"`
	Description string `json:"description"`
	Status string `json:"status"`
	Priority string `json:"priority"`
}