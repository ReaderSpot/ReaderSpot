package main

import (
	"autonomo_dos/controllers"
	"autonomo_dos/db"
	"autonomo_dos/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := db.ConectarDB()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(models.Libro{})
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/libro/{id}", controllers.Enrutamiento).Methods("PUT")
	router.HandleFunc("/libro", controllers.Enrutamiento).Methods("DELETE")
	router.HandleFunc("/libro", controllers.Enrutamiento).Methods("GET")
	router.HandleFunc("/libro", controllers.CrearLibro).Methods("POST")
	http.ListenAndServe(":5000", router)
}

/*
func main() {
	DB, err := db.ConectarDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer DB.Close()

	fmt.Println("Conexión exitosa a la base de datos MySQL.")
}
*/
