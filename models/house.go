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

type TabOverview struct {
	ID                int    `gorm:"primary_key; column:id"`
	CityCode          string `gorm:"column:city_code"`
	Date              string `gorm:"column:date"`
	NewHouseNum       int    `gorm:"column:new_house_num"`
	NewPeopleNum      int    `gorm:"column:new_people_num"`
	VisitNum          int    `gorm:"column:visit_num"`
	SoldNumInNinety   int    `gorm:"column:sold_num_in_ninety"`
	OnsaleNumCurrent  int    `gorm:"column:onsale_num_current"`
	AvgPriceLastMonth int    `gorm:"column:avg_price_last_month"`
}

func (TabOverview) TableName() string {
	return "TAB_OVERVIEW"
}

type TabTransRecords struct {
	ID         int    `gorm:"primary_key; column:id"`
	QuCode     string `gorm:"column:qu_code"`
	Date       string `gorm:"column:date"`
	HouseName  string `gorm:"column:house_name"`
	HouseCode  string `gorm:"column:house_code"`
	TotalPrice int    `gorm:"column:total_price"`
	Link       string `gorm:"column:link"`
	AvgPrice   int    `gorm:"column:avg_price"`
	TxId       string `gorm:"column:tx_id"`
}

func (TabTransRecords) TableName() string {
	return "TAB_TRANS_RECORDS"
}
