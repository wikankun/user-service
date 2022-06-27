package entity

type Verification struct {
	ID int `gorm:"not null" json:"id"`
	UserID int `gorm:"not null" json:"user_id"`
  	User User `gorm:"foreignKey:UserID"`
	Code string `gorm:"type:varchar(100);not null" json:"code"`
}
