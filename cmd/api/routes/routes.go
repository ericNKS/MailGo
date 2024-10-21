package routes

import (
	"sendMail/internal/handler"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/", handler.StartCampaign)
}
