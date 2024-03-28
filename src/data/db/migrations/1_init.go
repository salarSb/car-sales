package migrations

import (
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/constants"
	"github.com/salarSb/car-sales/data/db"
	"github.com/salarSb/car-sales/data/models"
	"github.com/salarSb/car-sales/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := db.GetDb()
	createTables(database)
	createDefaultInformation(database)
}

func Down1() {

}

func createTables(database *gorm.DB) {
	var tables []interface{}
	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	roleUser := models.RoleUser{}
	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, roleUser, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Migration, "cannot create tables through migration", nil)
		return
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)
	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)
	u := models.User{
		Username:     constants.DefaultUserName,
		FirstName:    constants.AdminFirstName,
		LastName:     constants.AdminLastName,
		MobileNumber: constants.AdminMobileNumber,
		Email:        constants.AdminEmail,
	}
	pass := constants.AdminPassword
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	createAdminUserIfNotExists(database, &u, adminRole.Id)
}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.Model(&models.Role{}).Select("1").Where("name = ?", r.Name).First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.Model(&models.User{}).Select("1").Where("username = ?", u.Username).First(&exists)
	if exists == 0 {
		database.Create(u)
		ru := models.RoleUser{UserId: u.Id, RoleId: roleId}
		database.Create(&ru)
	}
}
