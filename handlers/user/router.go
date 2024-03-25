package user

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET(":user_id", obterUsuarioPeloID)
}
