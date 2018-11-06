package models

type HouseEntity struct {
	ID      int64  `gorm:"primary_key; column:id"`
	Code    string `gorm:"column:code"`
	Online  string `gorm:"column:online"`
	Title   string `gorm:"column:title"`
	Created string `gorm:"column:created"`
}

func (HouseEntity) TableName() string {
	return "TAB_HOUSE_DETAIL"
}
