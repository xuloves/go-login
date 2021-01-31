package model

type Member struct {
	Id           int64   `gorm:"primary_key AUTO_INCREMENT" json:"id"`
	UserName     string  `gorm:"type:varchar(20)" json:"user_name"`
	Mobile       string  `gorm:"type:varchar(11)" json:"mobile"`
	Password     string  `gorm:"type:varchar(255)" json:"password"`
	RegisterTime int64   `gorm:"type:bigint" json:"register_time"`
	Avatar       string  `gorm:"type:varchar(255)" json:"avatar"`
	Balance      float64 `gorm:"type:double" json:"balance"`
	IsActive     int8    `gorm:"type:tinyint" json:"is_active"`
	City         string  `gorm:"type:varchar(10)" json:"city"`
}

func (Member) TableName() string {
	return "member"
}
