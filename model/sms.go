package model

type Sms struct {
	Id         int64  `gorm:"AUTO_INCREMENT" json:"id"`
	Phone      string `gorm:"varchar(11)" json:"phone"`
	BizId      string `gorm:"varchar(30)" json:"biz_id"`
	Code       string `gorm:"varchar(4)" json:"code"`
	CreateTime int64  `gorm:"bigint" json:"create_time"`
}

func (Sms) TableName() string {
	return "sms"
}
