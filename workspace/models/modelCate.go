package models

type ModelCate struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Sort       *int   `gorm:"not null" json:"sort"`
	Name       string `gorm:"not null" json:"name"`
	Visible    *int   `gorm:"not null" json:"visible"`
	Bill       *int   `gorm:"not null" json:"bill"`
	Coin       *int   `gorm:"not null" json:"coin"`
	Frequency  *int   `gorm:"not null" json:"frequency"`
	Model      string `gorm:"not null" json:"model"`
	MaxToken   *int   `gorm:"not null" json:"max_token"`
	Upload     *int   `gorm:"not null" json:"upload"`
	UploadType string `gorm:"not null" json:"upload_type"`
	Net        *int   `gorm:"not null" json:"net"`
	Tool       *int   `gorm:"not null" json:"tool"`
	History    *int   `gorm:"not null" json:"history"`
}

type ModelCateIDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (ModelCate) TableName() string {
	return "model_cate"
}
