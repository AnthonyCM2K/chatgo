package db

import (
	"context"

	"github.com/AnthonyCM2K/chatgo/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx := context.TODO()
	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
