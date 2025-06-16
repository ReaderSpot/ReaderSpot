package controllers

import (
	"autonomo_dos/db"
	"autonomo_dos/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Enrutamiento(w http.ResponseWriter, r *http.Request) {
	log.Print("Enrutamiento> " + r.Method)
	switch r.Method {

	case http.MethodGet:
		log.Print("Metodo get")
		buscarLibro(w, r)

	case http.MethodPut:
		log.Print("Metodo put")
		actualizarLibro(w, r)

	case http.MethodDelete:
		log.Print("Metodo delete")
		eliminarLibro(w, r)

	default:

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Metodo no permitido: "))

	}

}

func buscarLibro(w http.ResponseWriter, r *http.Request) {
	log.Print("Metodo buscar")
	id := r.URL.Query().Get("id")
	log.Print(id)

	var libros []models.Libro
	if id == "" {
		db.DB.Find(&libros)
	} else {
		params := map[string]interface{}{
			"id": id,
		}
		db.DB.Where(params).Find(&libros)
	}
	json.NewEncoder(w).Encode(&libros)
}

func CrearLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro
	json.NewDecoder(r.Body).Decode(&libro)
	if libro.UrlPDF == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Debes asignar una URL"))
		return
	}
	libroAgregado := db.DB.Create(&libro)
	err := libroAgregado.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al guardar en BD: " + err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&libro)
}

func actualizarLibro(w http.ResponseWriter, r *http.Request) {

	var libro models.Libro
	//id := extraerID(r.URL.Path)

	vars := mux.Vars(r)
	id := vars["id"]
	params := map[string]interface{}{
		"id": id,
	}
	db.DB.Where(params).First(&libro)

	var libroEntrante models.Libro
	json.NewDecoder(r.Body).Decode(&libroEntrante)
	if libroEntrante.UrlPDF == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Debes asignar una URL"))
		return
	}

	libro.Titulo = libroEntrante.Titulo
	libro.Autor = libroEntrante.Autor
	libro.Leido = libroEntrante.Leido
	libro.UrlPDF = libroEntrante.UrlPDF

	libroActualizado := db.DB.Updates(&libro)
	err := libroActualizado.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al actualizar en BD: " + err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&libro)

}

func eliminarLibro(w http.ResponseWriter, r *http.Request) {
	var libro models.Libro
	id := r.URL.Query().Get("id")
	log.Print("ID en delete: " + id)
	db.DB.First(&libro, id)
	db.DB.Delete(&libro)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&libro)
}
