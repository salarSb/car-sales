package dto

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
	Id             int                     `json:"id"`
	Name           string                  `json:"name"`
	Gearbox        GearboxResponse         `json:"gearbox"`
	Company        CompanyResponse         `json:"company"`
	CarType        CarTypeResponse         `json:"carType"`
	CarModelColors []CarModelColorResponse `json:"carModelColors,omitempty"`
	CarModelYears  []CarModelYearResponse  `json:"carModelYears,omitempty"`
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
	Id         int                     `json:"id"`
	Year       YearWithoutDateResponse `json:"year,omitempty"`
	CarModelId int                     `json:"carModelId,omitempty"`
}
