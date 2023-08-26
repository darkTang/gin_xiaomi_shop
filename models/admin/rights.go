package admin

type Rights struct {
	Id          string   `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ModuleName  string   `json:"moduleName" gorm:"type:varchar(255);comment:模块名称"`
	Type        uint8    `json:"type"` // 节点类型：1 模块 2 菜单  3 操作
	ActionName  string   `json:"actionName" gorm:"varchar(255);comment:操作模块"`
	Url         string   `json:"url" gorm:"varchar(255)"`
	ModuleId    string   `json:"moduleId" gorm:"type:int unsigned"` // 此module_id 和 id 关联，module_id = 0表示顶级模块
	Sort        int      `json:"sort"`
	Description string   `json:"description" gorm:"varchar(255)"`
	AddTime     uint64   `json:"add_time" gorm:"autoCreateTime:milli"`
	Status      uint8    `json:"status" gorm:"default:1"`
	RightsItem  []Rights `json:"rightsList" gorm:"foreignKey:ModuleId;references:Id"` // 自关联
}
