package models

type ModelKeys struct {
	ID          int    `gorm:"primary_key" json:"id"`
	ModelCateID int    `gorm:"not null" json:"model_cate_id"`
	Key         string `gorm:"not null" json:"key"`
	ApiAddr     string `gorm:"not null" json:"api_addr"`
	Weight      int    `gorm:"not null" json:"weight"`
	Enable      *int   `gorm:"not null" json:"enable" `
}

type ModelKeyCate struct {
	Id            int    `json:"id"`
	ModelCateID   int    `json:"model_cate_id"`
	ModelCateName string `json:"model_cate_name"`
	Key           string `json:"key"`
	ApiAddr       string `json:"api_addr"`
	Enable        int    `json:"enable"`
	Weight        int    `json:"weight"`
	Tool          int    `json:"tool"`
}

func (ModelKeyCate) TableName() string {
	return "view_model_key_cate"
}

func (ModelKeys) TableName() string {
	return "model_key"
}
