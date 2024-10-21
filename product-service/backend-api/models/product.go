package models

type Product struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}
