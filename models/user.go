package models

type User struct {
	ID          int64  `gorm:"primary_key;auto_increment" json:"id"`
	NamaLengkap string `gorm:"varchar(255);not null" json:"nama_lengkap"`
	Username    string `gorm:"varchar(255);not null" json:"username"`
	Password    string `gorm:"varchar(255);not null" json:"password"`
}
