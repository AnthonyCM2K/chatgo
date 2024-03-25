package routers

import (
	"encoding/json"
	"fmt"

	"github.com/AnthonyCM2K/chatgo/db"
	"github.com/AnthonyCM2K/chatgo/models"
	"github.com/aws/aws-lambda-go/events"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	fmt.Println("Entré en Ver Perfil")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	perfil, err := db.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Ocurrió un error al intentar buscar el registro " + err.Error()
		return r
	}

	respJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos de los usuarios como JSON " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
