package jwt

import (
	"time"

	"github.com/KBJ-ID/RedS/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT genero el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("KBJMasterDelDesarrollo")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioWeb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
