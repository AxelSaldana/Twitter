package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://axel:123456789@twitter.acy6k.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func ConectarBD() *mongo.ClientOptions {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Printfln("Conexion exitosa con la BD")
	return client
}
