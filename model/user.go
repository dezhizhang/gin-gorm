package model

type User struct {
	ID       int    `json:"ID"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	DeleteAt string `json:"delete_at"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Sex      int    `json:"sex"`
}
