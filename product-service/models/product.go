package models

type Product struct {
	Id      int     `json:"id" gorm:"primary_key"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Image   string  `json:"image"`
	Price   float64 `json:"Price"`
	Stock	int		`json:"Stock"`
}