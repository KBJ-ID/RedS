package routers

import (
	"net/http"

	"github.com/KBJ-ID/RedS/bd"
)

/*EliminarTweet es la funcion para eliminar un tweet*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar un tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
