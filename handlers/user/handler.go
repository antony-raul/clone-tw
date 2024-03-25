package user

import (
	userApp "github.com/antony-raul/tw/application/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func obterUsuarioPeloID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user, err := userApp.ObterUsuarioPeloID(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
