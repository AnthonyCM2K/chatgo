package db

import (
	"context"

	"github.com/AnthonyCM2K/chatgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx := context.TODO()

	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	var result []models.DevuelvoTweetsSeguidores

	cursor, err := col.Aggregate(ctx, condiciones)
	if err != nil {
		return result, false
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
