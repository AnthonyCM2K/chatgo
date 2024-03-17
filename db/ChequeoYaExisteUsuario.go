package db

import (
	"context"

	"github.com/AnthonyCM2K/chatgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx := context.TODO()

	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("usuarios")

	condition := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID

}
