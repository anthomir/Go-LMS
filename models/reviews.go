package models

type Review struct {
	UUIDModel
	Course      Course  `gorm:"foreignKey:CourseId" json:"course"`
	Rating      float64 `gorm:"type:double precision" validate:"min=1,max=5" json:"rating"`
	Description string  `gorm:"type:text" json:"description"`
	CreatedBy   User    `gorm:"foreignKey:UserId" json:"createdBy"`
}

func (Review) TableName() string {
	return "review"
}
