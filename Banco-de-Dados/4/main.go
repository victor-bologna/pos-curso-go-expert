package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ManyToManyProduct struct {
	ID                   int `gorm:"primaryKey"`
	Name                 string
	Price                float64
	ManyToManyCategories []ManyToManyCategory `gorm:"many2many:products_categories;"`
	gorm.Model
}

type ManyToManyCategory struct {
	ID                 int `gorm:"primaryKey"`
	Name               string
	ManyToManyProducts []ManyToManyProduct `gorm:"many2many:products_categories;"`
}

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func loadConfig(fileName string) (Config, error) {
	var config Config
	configFile, err := os.Open(fileName)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func main() {
	config, err := loadConfig("../config.json")
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Transform Struct into table
	db.AutoMigrate(&ManyToManyCategory{}, &ManyToManyProduct{})

	// Create Many to Many category
	/*categories := []ManyToManyCategory{
		{Name: "Kitchen"},
		{Name: "Electronics"},
	}
	db.Save(&categories)

	// Create product
	product := ManyToManyProduct{
		Name:                 "Electric pan",
		Price:                300.20,
		ManyToManyCategories: []ManyToManyCategory{categories[0], categories[1]},
	}
	db.Save(&product)*/

	var many2manyCategories []ManyToManyCategory
	err = db.Model(&ManyToManyCategory{}).Preload("ManyToManyProducts").Find(&many2manyCategories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range many2manyCategories {
		println(category.Name, ":")
		for _, product := range category.ManyToManyProducts {
			println("-", product.Name)
		}
	}
}
