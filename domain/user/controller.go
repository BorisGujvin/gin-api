package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, GetUserList())
}

func Create(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.PassHash), 14)
	user.PassHash = string(hash)

	user, err := StoreUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return

	}
	ctx.JSON(http.StatusOK, user)
}
