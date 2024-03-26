package user

import (
	userApp "github.com/antony-raul/tw/application/user"
	userModel "github.com/antony-raul/tw/models/user"
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

	user, err := userApp.ObterUsuarioPeloID(ctx, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func criarUsuario(ctx *gin.Context) {
	var (
		req *userModel.ReqUsuario
		ID  int64
		err error
	)

	if err = ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if ID, err = userApp.CriarUsuario(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, ID)
}
