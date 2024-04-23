package dto

import "time"

type CarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=50"`
}

type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GearboxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=50"`
}

type GearboxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCompanyRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=3,max=20"`
	CountryId int    `json:"countryId" binding:"required"`
}

type UpdateCompanyRequest struct {
	Name      string `json:"name,omitempty" binding:"alpha,min=3,max=20"`
	CountryId int    `json:"countryId,omitempty"`
}

type CompanyResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

type CreateCarModelRequest struct {
	Name      string `json:"name" binding:"required,min=3,max=100"`
	CompanyId int    `json:"companyId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
}

type UpdateCarModelRequest struct {
	Name      string `json:"name" binding:"min=3,max=100"`
	CompanyId int    `json:"companyId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
}

type CarModelResponse struct {
	Id                 int                        `json:"id"`
	Name               string                     `json:"name"`
	Gearbox            GearboxResponse            `json:"gearbox"`
	Company            CompanyResponse            `json:"company"`
	CarType            CarTypeResponse            `json:"carType"`
	CarModelColors     []CarModelColorResponse    `json:"carModelColors,omitempty"`
	CarModelYears      []CarModelYearResponse     `json:"carModelYears,omitempty"`
	CarModelFiles      []CarModelFileResponse     `json:"carModelFiles,omitempty"`
	CarModelProperties []CarModelPropertyResponse `json:"carModelProperties,omitempty"`
	CarModelComments   []CarModelCommentResponse  `json:"carModelComments,omitempty"`
}

type CreateCarModelYearRequest struct {
	CarModelId int `json:"carModelId" binding:"required"`
	YearId     int `json:"yearId" binding:"required"`
}

type UpdateCarModelYearRequest struct {
	CarModelId int `json:"carModelId,omitempty"`
	YearId     int `json:"yearId,omitempty"`
}

type CarModelYearResponse struct {
	Id                     int                            `json:"id"`
	Year                   YearWithoutDateResponse        `json:"year,omitempty"`
	CarModelId             int                            `json:"carModelId,omitempty"`
	CarModelPriceHistories []CarModelPriceHistoryResponse `json:"carModelPriceHistories,omitempty"`
}

type CreateCarModelPriceHistoryRequest struct {
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	PriceAt        time.Time `json:"priceAt" binding:"required"`
	Price          float64   `json:"price" binding:"required"`
}

type UpdateCarModelPriceHistoryRequest struct {
	PriceAt time.Time `json:"priceAt,omitempty"`
	Price   float64   `json:"price,omitempty"`
}

type CarModelPriceHistoryResponse struct {
	Id             int       `json:"id"`
	CarModelYearId int       `json:"carModelYearId"`
	Price          float64   `json:"price"`
	PriceAt        time.Time `json:"priceAt"`
}

type CreateCarModelFileRequest struct {
	CarModelId int  `json:"carModelId" binding:"required"`
	FileId     int  `json:"fileId" binding:"required"`
	IsMainFile bool `json:"isMainFile"`
}

type UpdateCarModelFileRequest struct {
	IsMainFile bool `json:"isMainFile" binding:"required"`
}

type CarModelFileResponse struct {
	Id         int          `json:"id"`
	CarModelId int          `json:"carModelId,omitempty"`
	File       FileResponse `json:"file,omitempty"`
	IsMainFile bool         `json:"isMainFile"`
}

type CreateCarModelPropertyRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	PropertyId int    `json:"propertyId" binding:"required"`
	Value      string `json:"value" binding:"required,max=1000"`
}

type UpdateCarModelPropertyRequest struct {
	Value string `json:"value" binding:"required,max=1000"`
}

type CarModelPropertyResponse struct {
	Id         int              `json:"id"`
	CarModelId int              `json:"carModelId,omitempty"`
	Property   PropertyResponse `json:"property,omitempty"`
	Value      string           `json:"value"`
}

type CreateCarModelCommentRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	UserId     int    `json:"userId"`
	Message    string `json:"message" binding:"required,max=500"`
}

type UpdateCarModelCommentRequest struct {
	Message string `json:"message" binding:"required,max=500"`
}

type CarModelCommentResponse struct {
	Id         int          `json:"id"`
	CarModelId int          `json:"carModelId"`
	User       UserResponse `json:"user"`
	Message    string       `json:"message"`
}
