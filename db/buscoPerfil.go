package db

import (
	"context"

	"github.com/AnthonyCM2K/chatgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx := context.TODO()
	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = "" //mandamos cadena vacia y al detectarlo por el omitemty lo omite en del json
	if err != nil {
		return perfil, err
	}

	return perfil, nil

}
