package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-foreign-language-rest-api/model"
	"go-foreign-language-rest-api/util"
	"net/http"
	"strconv"
)

var NewLanguage models.Language

func CreateLanguage(w http.ResponseWriter, r *http.Request) {
	CreateLanguage := &models.Language{}

	utils.ParseBody(r, CreateLanguage)
	b := CreateLanguage.CreateLanguage()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLanguage(w http.ResponseWriter, r *http.Request) {
	newLanguages := models.GetAllLanguages()
	res, _ := json.Marshal(newLanguages)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLanguageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	LanguageId := vars["LanguageId"]
	ID, err := strconv.ParseInt(LanguageId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	LanguageDetails, _ := models.GetLanguageById(ID)
	res, _ := json.Marshal(LanguageDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	var updateLanguage = &models.Language{}
	utils.ParseBody(r, updateLanguage)
	vars := mux.Vars(r)
	LanguageId := vars["LanguageId"]
	ID, err := strconv.ParseInt(LanguageId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	LanguageDetails, db := models.GetLanguageById(ID)

	if updateLanguage.Name != "" {
		LanguageDetails.Name = updateLanguage.Name
	}

	if updateLanguage.ShortCode != "" {
		LanguageDetails.ShortCode = updateLanguage.ShortCode
	}

	db.Save(&LanguageDetails)
	res, _ := json.Marshal(LanguageDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	LanguageId := vars["LanguageId"]
	ID, err := strconv.ParseInt(LanguageId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	Language := models.DeleteLanguage(ID)
	res, _ := json.Marshal(Language)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
