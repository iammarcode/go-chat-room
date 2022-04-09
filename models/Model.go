package models

type Model struct {
	ID         int `gorm:"prmary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}


