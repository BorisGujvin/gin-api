package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func IsHealthy(ctx *gin.Context) {
    res := make(map[string]interface{}, 0)
    res["status"] = 200
    res["healthy"] = "OK"
    ctx.JSON(http.StatusOK, res)
}