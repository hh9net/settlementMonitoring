package types

//省外数据
type JieSuanWssj struct {
	Id      int64 `gorm:"AUTO_INCREMENT primary_key"` // 设置 id 为自增类型
	JiaoyId int64 `gorm:"unique;not null"`
	Jine    int64
}

//本省结算数据
type JieSuanJiangssj struct {
	Id      int64 `gorm:"AUTO_INCREMENT primary_key"` // 设置 id 为自增类型
	JiaoyId int64 `gorm:"unique;not null"`
	Jine    int64
}
