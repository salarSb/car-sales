package dto

import (
	"mime/multipart"
	"time"
)

type CountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}

type CountryResponse struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	Cities    []CityResponse    `json:"cities,omitempty"`
	Companies []CompanyResponse `json:"companies,omitempty"`
}

type CreateCityRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=3,max=20"`
	CountryId int    `json:"countryId" binding:"required"`
}

type UpdateCityRequest struct {
	Name      string `json:"name,omitempty" binding:"alpha,min=3,max=20"`
	CountryId int    `json:"countryId,omitempty"`
}

type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

type FileFormRequest struct {
	File *multipart.FileHeader `json:"file" form:"file" binding:"required" swaggerignore:"true"`
}

type UploadFileRequest struct {
	FileFormRequest
	Description string `json:"description" form:"description" binding:"required"`
}

type CreateFileRequest struct {
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MimeType    string `json:"mimeType"`
}

type UpdateFileRequest struct {
	Description string `json:"description" binding:"required"`
}

type FileResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MimeType    string `json:"mimeType"`
}

type CreateColorRequest struct {
	Name    string `json:"name" binding:"required,max=15"`
	HexCode string `json:"hexCode" binding:"required,max=7"`
}

type UpdateColorRequest struct {
	Name    string `json:"name,omitempty" binding:"max=15"`
	HexCode string `json:"hexCode,omitempty" binding:"max=7"`
}

type ColorResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	HexCode string `json:"hexCode,omitempty"`
}

type CreateCarModelColorRequest struct {
	CarModelId int `json:"carModelId" binding:"required"`
	ColorId    int `json:"colorId" binding:"required"`
}

type UpdateCarModelColorRequest struct {
	CarModelId int `json:"carModelId,omitempty"`
	ColorId    int `json:"colorId,omitempty"`
}

type CarModelColorResponse struct {
	Id    int           `json:"id"`
	Color ColorResponse `json:"color,omitempty"`
}

type CreateYearRequest struct {
	Title   string    `json:"title" binding:"required,min=4,max=4"`
	Year    int       `json:"year" binding:"required"`
	StartAt time.Time `json:"startAt" binding:"required"`
	EndAt   time.Time `json:"endAt" binding:"required"`
}

type UpdateYearRequest struct {
	Title   string    `json:"title,omitempty" binding:"min=4,max=4"`
	Year    int       `json:"year,omitempty"`
	StartAt time.Time `json:"startAt,omitempty"`
	EndAt   time.Time `json:"endAt,omitempty"`
}

type YearResponse struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Year    int       `json:"year"`
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
}

type YearWithoutDateResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}
