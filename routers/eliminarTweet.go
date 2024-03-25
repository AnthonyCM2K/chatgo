package routers

import (
	"github.com/aws/aws-lambda-go/events"

	"github.com/AnthonyCM2K/chatgo/db"
	"github.com/AnthonyCM2K/chatgo/models"
)

func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	err := db.BorroTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrió un error al intentar borrar el tweet " + err.Error()
		return r
	}

	r.Message = "Eliminar Tweet OK!"
	r.Status = 200
	return r
}
