package routes

import (
	"github.com/gorilla/mux"
	"go-foreign-language-rest-api/controller"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/languages", controllers.CreateLanguage).Methods("POST")
	router.HandleFunc("/languages", controllers.GetLanguage).Methods("GET")
	router.HandleFunc("/languages/{languageId}", controllers.GetLanguageById).Methods("GET")
	router.HandleFunc("/languages/{languageId}", controllers.UpdateLanguage).Methods("PUT")
	router.HandleFunc("/languages/{languageId}", controllers.DeleteLanguage).Methods("DELETE")
}
