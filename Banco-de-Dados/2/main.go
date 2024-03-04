package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductORM struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
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
	db.AutoMigrate(&ProductORM{})

	// Create one product
	/*db.Create(&ProductORM{
		Name:  "Notebook",
		Price: 1000.00,
	})

	// Create many products
	db.Create(&[]ProductORM{
		{Name: "Thermal paste", Price: 100.00},
		{Name: "Monitor 24 inches", Price: 4300.00},
		{Name: "Test", Price: 30.00},
	})*/

	//Find product By Id
	/*var productORM ProductORM
	db.Find(&productORM, 2)
	fmt.Println(productORM)
	productORM = ProductORM{}
	db.First(&productORM, "name = ?", "Thermal paste")
	fmt.Println(productORM)

	// Select all
	var productsORM []ProductORM
	db.Find(&productsORM)
	for _, product := range productsORM {
		fmt.Println(product)
	}

	println("\n----------------\n")
	// Limit search paged by offset (get page 2)
	productsORM = nil
	db.Limit(2).Offset(2).Find(&productsORM)
	for _, product := range productsORM {
		fmt.Println(product)
	}

	// all products above 100
	productsORM = nil
	db.Where("price > 100").Find(&productsORM)
	for _, product := range productsORM {
		fmt.Println(product)
	}

	// Like search
	productsORM = nil
	db.Where("name LIKE ?", "%book%").Find(&productsORM)
	for _, product := range productsORM {
		fmt.Println(product)
	}
	println("\n----------------\n")

	// Update name and checking new name.
	var p ProductORM
	db.First(&p, 1)
	p.Name = "New Notebook"
	db.Save(&p)*/

	var p ProductORM
	db.First(&p, 1)
	fmt.Println(p)

	// Delete product
	db.Delete(&p)
}
