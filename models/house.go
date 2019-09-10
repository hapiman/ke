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

type TabXiaoQuOverview struct {
	ID               int    `gorm:"primary_key; column:id"`
	CityCode         string `gorm:"column:city_code"`
	DistrictCode     string `gorm:"column:district_code"`
	AreaCode         string `gorm:"column:area_code"`
	QuCode           string `gorm:"column:qu_code"`
	Date             string `gorm:"column:date"`
	Name             string `gorm:"column:name"`
	OnsaleNumCurrent int    `gorm:"column:onsale_num_current"`
	AvgPrice         int    `gorm:"column:avg_price"`
	SoldNumInNinety  int    `gorm:"column:sold_num_in_ninety"`
	SoldNumInThirty  int    `gorm:"column:sold_num_in_thirty"`
	VisitNumInThirty int    `gorm:"column:visit_num_in_thirty"`
}

func (TabXiaoQuOverview) TableName() string {
	return "TAB_XIAOQU_OVERVIEW"
}
