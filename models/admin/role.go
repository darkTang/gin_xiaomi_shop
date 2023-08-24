package admin

type Role struct {
	Id          string `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement"`
	Title       string `json:"title" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	Status      uint8  `json:"status" gorm:"default:1"`
	AddTime     uint64 `json:"add_time" gorm:"autoCreateTime:milli"`
}

func (Role) TableName() string {
	return "role"
}
