package models

import "time"

type Gearbox struct {
	BaseModel
	Name      string `gorm:"size:50;type:string;not null;unique"`
	CarModels []CarModel
}

type CarType struct {
	BaseModel
	Name      string `gorm:"size:50;type:string;not null;unique"`
	CarModels []CarModel
}

type Company struct {
	BaseModel
	Name      string  `gorm:"size:50;type:string;not null;unique"`
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CountryId int
	CarModels []CarModel
}

type CarModel struct {
	BaseModel
	Name               string  `gorm:"size:100;type:string;not null;unique"`
	Company            Company `gorm:"foreignKey:CompanyId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CompanyId          int
	CarType            CarType `gorm:"foreignKey:CarTypeId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarTypeId          int
	Gearbox            Gearbox `gorm:"foreignKey:GearboxId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	GearboxId          int
	CarModelColors     []CarModelColor
	CarModelYears      []CarModelYear
	CarModelProperties []CarModelProperty
	CarModelComments   []CarModelComment
	CarModelFiles      []CarModelFile
}

type CarModelColor struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId int
	Color      Color `gorm:"foreignKey:ColorId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	ColorId    int
}

type CarModelYear struct {
	BaseModel
	CarModel               CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId             int      `gorm:"uniqueIndex:idx_CarModelId_YearId"`
	Year                   Year     `gorm:"foreignKey:YearId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	YearId                 int      `gorm:"uniqueIndex:idx_CarModelId_YearId"`
	CarModelPriceHistories []CarModelPriceHistory
}

type CarModelFile struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId int
	File       File `gorm:"foreignKey:FileId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	FileId     int
	IsMainFile bool
}

type CarModelPriceHistory struct {
	BaseModel
	CarModelYear   CarModelYear `gorm:"foreignKey:CarModelYearId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelYearId int
	Price          float64   `gorm:"type:decimal(10,2);not null"`
	PriceAt        time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
}

type CarModelProperty struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId int
	Property   Property `gorm:"foreignKey:PropertyId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	PropertyId int
	Value      string `gorm:"size:100;type:string;not null"`
}

type CarModelComment struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId int
	User       User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId     int
	Message    string `gorm:"size:500;type:string;not null"`
}
