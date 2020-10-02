package routers

import (
	"net/http"

	"github.com/KBJ-ID/RedS/bd"
	"github.com/KBJ-ID/RedS/models"
)

/*BajaRelacion realizar el borrado de la relación entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar relación"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relación"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
