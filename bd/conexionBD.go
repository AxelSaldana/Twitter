package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://axel:123456789@twitter.acy6k.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conecetar la BD*/
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
	log.Println("Conexion exitosa con la BD")
	return client
}

/*Chequeo connection es el ping de la BD */
func ChequeoConnection() int {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
