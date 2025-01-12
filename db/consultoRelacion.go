package db

import (
	"context"

	"github.com/AnthonyCM2K/chatgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(t models.Relacion) bool {
	ctx := context.TODO()
	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false
	}
	return true
}
