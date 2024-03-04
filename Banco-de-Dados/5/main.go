package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	tx := db.Begin()
	var category ManyToManyCategory
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, 1).Error //Lock Table for everyone
	if err != nil {
		panic(err)
	}
	category.Name = "New Electronics"
	tx.Debug().Save(&category) // Alter data
	tx.Commit()                // Commit and release line for everyone
}
