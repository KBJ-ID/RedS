package main

import (
	"log"

	"github.com/KBJ-ID/RedS/bd"
	"github.com/KBJ-ID/RedS/handlers"
)

/* Funcion Main sss*/
func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexión a la BD")
		return
	}
	handlers.Manejadores()
}
