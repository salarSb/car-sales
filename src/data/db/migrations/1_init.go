package migrations

import (
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/constants"
	"github.com/salarSb/car-sales/data/db"
	"github.com/salarSb/car-sales/data/models"
	"github.com/salarSb/car-sales/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := db.GetDb()
	createTables(database)
	createDefaultInformation(database)
	createCountry(database)
	createPropertyCategory(database)
	createCarType(database)
	createGearbox(database)
	createColor(database)
	createYear(database)
}

func Down1() {

}

func createTables(database *gorm.DB) {
	var tables []interface{}

	// Basic
	tables = addNewTable(database, models.Country{}, tables)
	tables = addNewTable(database, models.City{}, tables)
	tables = addNewTable(database, models.File{}, tables)
	tables = addNewTable(database, models.Year{}, tables)
	// Property
	tables = addNewTable(database, models.PropertyCategory{}, tables)
	tables = addNewTable(database, models.Property{}, tables)

	// User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.RoleUser{}, tables)

	// Car
	tables = addNewTable(database, models.Company{}, tables)
	tables = addNewTable(database, models.Gearbox{}, tables)
	tables = addNewTable(database, models.Color{}, tables)
	tables = addNewTable(database, models.CarType{}, tables)

	tables = addNewTable(database, models.CarModel{}, tables)
	tables = addNewTable(database, models.CarModelColor{}, tables)
	tables = addNewTable(database, models.CarModelYear{}, tables)
	tables = addNewTable(database, models.CarModelFile{}, tables)
	tables = addNewTable(database, models.CarModelPriceHistory{}, tables)
	tables = addNewTable(database, models.CarModelProperty{}, tables)
	tables = addNewTable(database, models.CarModelComment{}, tables)

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

func createCountry(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Country{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}, Companies: []models.Company{
			{Name: "Saipa"},
			{Name: "Iran khodro"},
		}})
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}, Companies: []models.Company{
			{Name: "Tesla"},
			{Name: "Jeep"},
		}})
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}, Companies: []models.Company{
			{Name: "Opel"},
			{Name: "Benz"},
		}})
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}, Companies: []models.Company{
			{Name: "Chery"},
			{Name: "Geely"},
		}})
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}, Companies: []models.Company{
			{Name: "Ferrari"},
			{Name: "Fiat"},
		}})
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}, Companies: []models.Company{
			{Name: "Renault"},
			{Name: "Bugatti"},
		}})
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}, Companies: []models.Company{
			{Name: "Toyota"},
			{Name: "Honda"},
		}})
		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}, Companies: []models.Company{
			{Name: "Kia"},
			{Name: "Hyundai"},
		}})
	}
}

func createPropertyCategory(database *gorm.DB) {
	count := 0

	database.
		Model(&models.PropertyCategory{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.PropertyCategory{Name: "Body"})                     // بدنه
		database.Create(&models.PropertyCategory{Name: "Engine"})                   // موتور
		database.Create(&models.PropertyCategory{Name: "Drivetrain"})               // پیشرانه
		database.Create(&models.PropertyCategory{Name: "Suspension"})               // تعلیق
		database.Create(&models.PropertyCategory{Name: "Equipment"})                // تجهیزات
		database.Create(&models.PropertyCategory{Name: "Driver support systems"})   // سیستم های پشتیبانی راننده
		database.Create(&models.PropertyCategory{Name: "Lights"})                   // چراغ ها
		database.Create(&models.PropertyCategory{Name: "Multimedia"})               // چند رسانه ای
		database.Create(&models.PropertyCategory{Name: "Safety equipment"})         // تجهیزات ایمنی
		database.Create(&models.PropertyCategory{Name: "Seats and steering wheel"}) // صندلی و فرمان
		database.Create(&models.PropertyCategory{Name: "Windows and mirrors"})      // پنجره و آینه
	}
	createProperty(database, "Body")
	createProperty(database, "Engine")
	createProperty(database, "Drivetrain")
	createProperty(database, "Suspension")
	createProperty(database, "Comfort")
	createProperty(database, "Driver support systems")
	createProperty(database, "Lights")
	createProperty(database, "Multimedia")
	createProperty(database, "Safety equipment")
	createProperty(database, "Seats and steering wheel")
	createProperty(database, "Windows and mirrors")

}

func createProperty(database *gorm.DB, cat string) {
	count := 0
	catModel := models.PropertyCategory{}

	database.
		Model(models.PropertyCategory{}).
		Where("name = ?", cat).
		Find(&catModel)

	database.
		Model(&models.Property{}).
		Select("count(*)").
		Where("property_category_id = ?", catModel.Id).
		Find(&count)

	if count > 0 || catModel.Id == 0 {
		return
	}
	var props *[]models.Property
	switch cat {
	case "Body":
		props = getBodyProperties(catModel.Id)

	case "Engine":
		props = getEngineProperties(catModel.Id)

	case "Drivetrain":
		props = getDrivetrainProperties(catModel.Id)

	case "Suspension":
		props = getSuspensionProperties(catModel.Id)

	case "Comfort":
		props = getComfortProperties(catModel.Id)

	case "Driver support systems":
		props = getDriverSupportSystemProperties(catModel.Id)

	case "Lights":
		props = getLightsProperties(catModel.Id)

	case "Multimedia":
		props = getMultimediaProperties(catModel.Id)

	case "Safety equipment":
		props = getSafetyEquipmentProperties(catModel.Id)

	case "Seats and steering wheel":
		props = getSeatsProperties(catModel.Id)

	case "Windows and mirrors":
		props = getWindowsProperties(catModel.Id)

	default:
		props = &([]models.Property{})
	}

	for _, prop := range *props {
		database.Create(&prop)
	}
}

func createCarType(database *gorm.DB) {
	count := 0
	database.
		Model(&models.CarType{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.CarType{Name: "Crossover"})
		database.Create(&models.CarType{Name: "Sedan"})
		database.Create(&models.CarType{Name: "Sports"})
		database.Create(&models.CarType{Name: "Coupe"})
		database.Create(&models.CarType{Name: "Hatchback"})
	}
}

func createGearbox(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Gearbox{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Gearbox{Name: "Manual"})
		database.Create(&models.Gearbox{Name: "Automatic"})
	}
}

func createColor(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Color{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Color{Name: "Black", HexCode: "#000000"})
		database.Create(&models.Color{Name: "White", HexCode: "#ffffff"})
		database.Create(&models.Color{Name: "Blue", HexCode: "#0000ff"})
	}
}

func createYear(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Year{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Year{
			Title:   "1402",
			Year:    1402,
			StartAt: time.Date(2023, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:   time.Date(2024, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.Year{
			Title:   "1401",
			Year:    1401,
			StartAt: time.Date(2022, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:   time.Date(2023, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.Year{
			Title:   "1400",
			Year:    1400,
			StartAt: time.Date(2021, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:   time.Date(2022, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.Year{
			Title:   "1399",
			Year:    1399,
			StartAt: time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
			EndAt:   time.Date(2021, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.Year{
			Title:   "1398",
			Year:    1398,
			StartAt: time.Date(2019, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:   time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC),
		})

		database.Create(&models.Year{
			Title:   "1398",
			Year:    1398,
			StartAt: time.Date(2018, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
			EndAt:   time.Date(2019, time.Month(3), 21, 0, 0, 0, 0, time.UTC),
		})
	}
}
