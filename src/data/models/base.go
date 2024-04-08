package models

type Country struct {
	BaseModel
	Name   string `gorm:"size:15;type:string;not null;unique"`
	Cities []City
}

type City struct {
	BaseModel
	Name      string `gorm:"size:10;type:string;not null;unique"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
}
