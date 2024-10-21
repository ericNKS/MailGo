package handler

import (
	"net/http"
	"sendMail/internal/model"
	"sendMail/internal/service"
	"sync"

	"github.com/gin-gonic/gin"
)

func StartCampaign(ctx *gin.Context) {
	var body model.Campaign

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cp, err := service.CreateCampaign(&body.Remetente, body.Destinatarios, &body.Mensagem)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go cp.Execute(&wg)
	wg.Wait()
	ctx.JSON(http.StatusAccepted, gin.H{
		"success": cp,
	})
}
