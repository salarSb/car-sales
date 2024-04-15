package models

type PropertyCategory struct {
	BaseModel
	Name       string     `gorm:"size:100;type:string;not null"`
	Icon       string     `gorm:"size:1000;type:string;null"`
	Properties []Property `gorm:"foreignKey:PropertyCategoryId"`
}

type Property struct {
	BaseModel
	Name               string           `gorm:"size:100;type:string;not null"`
	Icon               string           `gorm:"size:1000;type:string;null"`
	PropertyCategory   PropertyCategory `gorm:"foreignKey:PropertyCategoryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	PropertyCategoryId int
	Description        string `gorm:"size:1000;type:string;null"`
	DataType           string `gorm:"size:15;type:string;not null"`
	Unit               string `gorm:"size:15;type:string;not null"`
	CarModelProperties []CarModelProperty
}
