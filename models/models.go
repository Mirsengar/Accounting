package models

type Product struct {
	ID           uint `gorm:"primarykey;autoIncrement"`
	CategoryName string
	Category     Category `gorm:"foreignKey:CategoryName"`
	Name         string   `gorm:"unique;not null"`
	Price        float64
}
type Category struct {
	ID   uint   `gorm:"primarykey;autoIncrement"`
	Name string `gorm:"unique;not null"`
}
type Cart struct {
	ID           uint `gorm:"primarykey;autoIncrement"`
	UserID       uint
	User         User `gorm:"foreignKey:UserID"`
	ProductID    uint
	CartProducts []Product `gorm:"foreignKey:ID;references:ProductID"`
}
type Invoice struct {
	ID          uint `gorm:"primarykey;autoIncrement"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	TotalAmount float64
	Discount    float64
	FinalAmount float64
}
type User struct {
	ID      uint   `gorm:"primarykey;autoIncrement"`
	Name    string `gorm:"unique;not null" json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}
