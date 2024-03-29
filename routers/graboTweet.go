package routers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/AnthonyCM2K/chatgo/db"
	"github.com/AnthonyCM2K/chatgo/models"
)

func GraboTweet(ctx context.Context, claim models.Claim) models.RespApi {
	var mensaje models.Tweet
	var r models.RespApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)
	if err != nil {
		r.Message = "Ocurrió un error al intentar decodificar el body " + err.Error()
		return r
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertoTweet(registro)
	if err != nil {
		r.Message = "Ocurrió un error al intentar insertar el registro " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el Tweet"
		return r
	}

	r.Status = 200
	r.Message = "Tweet creado Correctamente"
	return r
}
