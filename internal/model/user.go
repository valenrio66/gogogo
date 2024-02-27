package model

type User struct {
	IdUser   int    `gorm:"primaryKey;column:id_user" json:"-"`
	NamaUser string `gorm:"column:nama_user" json:"nama_user"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}
