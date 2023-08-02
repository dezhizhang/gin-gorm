package model

type User struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}
