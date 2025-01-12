package routers

import (
	"context"

	"github.com/AnthonyCM2K/chatgo/db"
	"github.com/AnthonyCM2K/chatgo/models"
	"github.com/aws/aws-lambda-go/events"
)

func AltaRelacion(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parametro ID es Obligatorio"
		return r
	}

	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	status, err := db.InsertoRelacion(t)
	if err != nil {
		r.Message = "Ocurrio un error al intentar insertar relaciòn " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar la relación"
		return r
	}

	r.Status = 200
	r.Message = "Alta de Relacion OK"
	return r

}
