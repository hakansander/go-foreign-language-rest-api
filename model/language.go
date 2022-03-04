package models

import (
	"github.com/jinzhu/gorm"
	"go-foreign-language-rest-api/config"
)

var db *gorm.DB

type Language struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShortCode string `json:"shortCode"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Language{})
}

func (language *Language) CreateLanguage() *Language {
	db.NewRecord(language)
	db.Create(&language)
	return language
}

func GetAllLanguages() []Language {
	var Languages []Language
	db.Find(&Languages)
	return Languages
}

func GetLanguageById(Id int64) (*Language, *gorm.DB) {
	var getLanguage Language
	db := db.Where("ID = ?", Id).Find(&getLanguage)
	return &getLanguage, db
}

func DeleteLanguage(ID int64) Language {
	var language Language
	db.Where("ID = ?", ID).Delete(language)
	return language
}
