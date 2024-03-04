package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BelongsProduct struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryId   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct { // Product Belongs to (Category)
	ID              int `gorm:"primaryKey"`
	Name            string
	BelongsProducts []BelongsProduct // Category has Many (Products)
}

type SerialNumber struct { // Product Has One (Serial Number)
	ID               int `gorm:"primaryKey"`
	Number           string
	BelongsProductID int
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
	db.AutoMigrate(&BelongsProduct{}, &Category{}, &SerialNumber{})

	// Create category
	/*category := Category{
		Name: "Kitchen",
	}
	db.Save(&category)

	// Create product
	product := BelongsProduct{
		Name:       "Pan",
		Price:      300.20,
		CategoryId: category.ID,
	}
	db.Save(&product)

	// Create Serial Number
	serialNumber := SerialNumber{
		Number:           "123456",
		BelongsProductID: 2,
	}
	db.Save(&serialNumber)*/

	var products []BelongsProduct
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

	var categories []Category
	err = db.Model(&Category{}).Preload("BelongsProducts").Preload("BelongsProducts.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name, ":")
		for _, product := range category.BelongsProducts {
			println("-", product.Name, "S/N:", product.SerialNumber.Number)
		}
	}
}
