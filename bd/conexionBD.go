package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la base datos */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://user1:ShuuNu7BMVWP1Caw@cluster0.7i0ao.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConectarBD es la función que me permite conectar la Base de datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

/*ChequeoConnection es el ping a la base de datos */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
