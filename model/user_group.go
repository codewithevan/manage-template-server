package model

type UserGroup struct {
	Id         int    `gorm:"id" json:"id"`
	GroupName  string `gorm:"group_name" json:"group_name" `
	System     bool   `gorm:"system" json:"system"`
	CreateTime int    `gorm:"create_time" json:"create_time"`
	UpdateTime int    `gorm:"update_time" json:"update_time"`
}

type UserGroupPermission struct {
	Id      int `gorm:"id" json:"id"`
	GroupId int `gorm:"group_id" json:"group_id"`
	MenuId  int `gorm:"menu_id" json:"menu_id"`
}
