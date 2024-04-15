package migrations

import "github.com/salarSb/car-sales/data/models"

func getBodyProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Height",
			PropertyCategoryId: cat,
			Description:        "Height of the car",
			DataType:           "int",
			Unit:               "mm",
		}, {
			Name:               "Width",
			PropertyCategoryId: cat,
			Description:        "Width of the car",
			DataType:           "int",
			Unit:               "mm",
		}, {
			Name:               "Length",
			PropertyCategoryId: cat,
			Description:        "Length of the car",
			DataType:           "int",
			Unit:               "mm",
		},
	}

	return &props
}

func getEngineProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Displacement",
			PropertyCategoryId: cat,
			Description:        "Displacement of the car",
			DataType:           "int",
			Unit:               "cc",
		}, {
			Name:               "Engine power ",
			PropertyCategoryId: cat,
			Description:        "Engine power of the car",
			DataType:           "int",
			Unit:               "hp/kw/rpm",
		}, {
			Name:               "Maximum torque",
			PropertyCategoryId: cat,
			Description:        "Maximum torque of the car",
			DataType:           "int",
			Unit:               "nm",
		}, {
			Name:               "Fuel consumption",
			PropertyCategoryId: cat,
			Description:        "Fuel consumption (l/100km) (city/suburb/general) of the car",
			DataType:           "int",
			Unit:               "nm",
		}, {
			Name:               "Maximum speed",
			PropertyCategoryId: cat,
			Description:        "Maximum speed of the car",
			DataType:           "int",
			Unit:               "km/h",
		},
	}

	return &props
}

func getDrivetrainProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Gears",
			PropertyCategoryId: cat,
			Description:        "Gears of the car",
			DataType:           "int",
			Unit:               "count",
		},
	}

	return &props
}

func getSuspensionProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Front brakes",
			PropertyCategoryId: cat,
			Description:        "Front brakes of the car",
			DataType:           "string",
			Unit:               "",
		}, {
			Name:               "Rear brakes",
			PropertyCategoryId: cat,
			Description:        "Rear brakes of the car",
			DataType:           "string",
			Unit:               "",
		}, {
			Name:               "Tires",
			PropertyCategoryId: cat,
			Description:        "Tires of the car",
			DataType:           "string",
			Unit:               "",
		},
	}

	return &props
}

func getComfortProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Central lock",
			PropertyCategoryId: cat,
			Description:        "Central lock of the car",
			DataType:           "bool",
			Unit:               "",
		}, {
			Name:               "Remote lock",
			PropertyCategoryId: cat,
			Description:        "Remote lock of the car",
			DataType:           "bool",
			Unit:               "",
		}, {
			Name:               "Keyless engine start",
			PropertyCategoryId: cat,
			Description:        "Keyless engine start of the car",
			DataType:           "bool",
			Unit:               "",
		},
	}

	return &props
}

func getDriverSupportSystemProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Parking radar",
			PropertyCategoryId: cat,
			Description:        "Parking radar of the car",
			DataType:           "bool",
			Unit:               "",
		}, {
			Name:               "Anti-lock braking system (ABS)",
			PropertyCategoryId: cat,
			Description:        "Anti-lock braking system (ABS) of the car",
			DataType:           "bool",
			Unit:               "",
		}, {
			Name:               "Electronic brakeforce distribution (EBD)",
			PropertyCategoryId: cat,
			Description:        "Electronic brakeforce distribution (EBD) of the car",
			DataType:           "bool",
			Unit:               "",
		}, {
			Name:               "Brake Assist",
			PropertyCategoryId: cat,
			Description:        "Brake Assist of the car",
			DataType:           "bool",
			Unit:               "",
		},
	}

	return &props
}

func getLightsProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Low beam",
			PropertyCategoryId: cat,
			Description:        "Low beam of the car",
			DataType:           "string",
			Unit:               "",
		}, {
			Name:               "High beam",
			PropertyCategoryId: cat,
			Description:        "High beam of the car",
			DataType:           "string",
			Unit:               "",
		}, {
			Name:               "Daytime running lights",
			PropertyCategoryId: cat,
			Description:        "Daytime running lights of the car",
			DataType:           "bool",
			Unit:               "",
		}, {
			Name:               "Automatic headlights",
			PropertyCategoryId: cat,
			Description:        "Automatic headlights of the car",
			DataType:           "bool",
			Unit:               "",
		},
	}

	return &props
}

func getMultimediaProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Bluetooth",
			PropertyCategoryId: cat,
			Description:        "Bluetooth of the car",
			DataType:           "bool",
			Unit:               "",
		},
	}

	return &props
}

func getSafetyEquipmentProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Driver front airbag",
			PropertyCategoryId: cat,
			Description:        "Driver front airbag of the car",
			DataType:           "bool",
			Unit:               "",
		}, {
			Name:               "Passenger front airbag",
			PropertyCategoryId: cat,
			Description:        "Passenger front airbag of the car",
			DataType:           "bool",
			Unit:               "",
		},
	}

	return &props
}

func getSeatsProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Driver seat electric adjustment",
			PropertyCategoryId: cat,
			Description:        "Driver seat electric adjustment of the car",
			DataType:           "bool",
			Unit:               "",
		}, {
			Name:               "Passenger seat electric adjustment",
			PropertyCategoryId: cat,
			Description:        "Passenger seat electric adjustment of the car",
			DataType:           "bool",
			Unit:               "",
		},
	}

	return &props
}

func getWindowsProperties(cat int) *[]models.Property {
	var props = []models.Property{
		{
			Name:               "Power windows",
			PropertyCategoryId: cat,
			Description:        "Power windows of the car",
			DataType:           "bool",
			Unit:               "",
		}, {
			Name:               "Rear wiper",
			PropertyCategoryId: cat,
			Description:        "Rear wiper of the car",
			DataType:           "bool",
			Unit:               "",
		},
	}

	return &props
}
