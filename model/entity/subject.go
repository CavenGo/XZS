package entity

type Subject struct {
	Id        int    `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`
	Level     int    `json:"level" gorm:"column:level"`
	LevelName string `json:"levelName" gorm:"column:level_name"`
	ItemOrder int    `json:"itemOrder" gorm:"column:item_order"`
	Deleted   bool   `json:"deleted" gorm:"column:deleted"`
}
