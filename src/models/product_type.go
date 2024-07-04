package models

type ProductType struct {
	ID       int       `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Products []Product `json:",omitempty" gorm:"foreignKey:TypeId"`
}
