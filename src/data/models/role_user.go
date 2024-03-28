package models

type RoleUser struct {
	BaseModel
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	RoleId int
	UserId int
}
