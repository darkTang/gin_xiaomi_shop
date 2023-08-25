package admin

type Manager struct {
	Id       string `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Mobile   string `json:"mobile" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Status   uint8  `json:"status" gorm:"default:1"`
	RoleId   string `json:"role_id" gorm:"type:tinyint unsigned"`
	AddTime  uint64 `json:"add_time" gorm:"autoCreateTime:milli"`
	IsSuper  uint8  `json:"is_super" gorm:"default:0"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleId;references:Id"`
}

func (Manager) TableName() string {
	return "manager"
}
