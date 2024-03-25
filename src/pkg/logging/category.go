package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Internal        Category = "Internal"
	Postgres        Category = "Category"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
)

const (
	// StartUp General
	StartUp SubCategory = "StartUp"
	// ExternalService General
	ExternalService SubCategory = "ExternalService"

	// Select Postgres
	Select SubCategory = "Select"
	// Rollback Postgres
	Rollback SubCategory = "Rollback"
	// Update Postgres
	Update SubCategory = "Update"
	// Delete Postgres
	Delete SubCategory = "Delete"
	// Insert Postgres
	Insert SubCategory = "Insert"

	// Api Internal
	Api SubCategory = "Api"
	// HashPassword Internal
	HashPassword SubCategory = "HashPassword"
	// DefaultRoleNotFound Internal
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"

	// MobileValidation  Validation
	MobileValidation SubCategory = "MobileValidation"
	// PasswordValidation  Validation
	PasswordValidation SubCategory = "PasswordValidation"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "Logger"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Body         ExtraKey = "Body"
	ErrorMessage ExtraKey = "ErrorMessage"
)
