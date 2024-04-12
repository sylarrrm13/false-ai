package models

type Codelkup struct {
	Listname string `gorm:"column:listname;primary_key" json:"listname"`
	Value    string `gorm:"column:value" json:"value"`
	Ref      string `gorm:"column:ref" json:"ref"`
	Ref2     string `gorm:"column:ref2" json:"ref2"`
}

func (Codelkup) TableName() string {
	return "codelkup"
}
