package constants

const (
	// AdminRoleName => User
	AdminRoleName string = "admin"
	// DefaultRoleName => User
	DefaultRoleName string = "default"
	// DefaultUserName => User
	DefaultUserName string = "admin"
	// AdminFirstName => User
	AdminFirstName string = "Salar"
	// AdminLastName => User
	AdminLastName string = "Sabeti"
	// AdminMobileNumber => User
	AdminMobileNumber string = "09999999999"
	// AdminPassword => User
	AdminPassword string = "12345678"
	// AdminEmail => User
	AdminEmail string = "admin@admin.com"
	// RedisOtpDefaultKey => User
	RedisOtpDefaultKey string = "otp"

	// AuthorizationHeaderKey => Claims
	AuthorizationHeaderKey string = "Authorization"
	// UserIdKey => Claims
	UserIdKey string = "user_id"
	// FirstNameKey => Claims
	FirstNameKey string = "first_name"
	// LastNameKey => Claims
	LastNameKey string = "last_name"
	// UsernameKey => Claims
	UsernameKey string = "username"
	// EmailKey => Claims
	EmailKey string = "email"
	// MobileNumberKey => Claims
	MobileNumberKey string = "mobile_number"
	// RolesKey => Claims
	RolesKey string = "roles"
	// ExpireTimeKey => Claims
	ExpireTimeKey string = "exp"
)
