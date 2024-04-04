package content

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetContents(ctx *gin.Context) {
    name, present := ctx.GetQuery("name")
    if !present {
        ctx.Error(fmt.Errorf("name is required"))
        ctx.AbortWithStatus(http.StatusBadRequest)
        return
    }
    response := make(map[string]string, 0)
    response["name"] = name

    ctx.JSON(http.StatusOK, response)
}

func PostContents(ctx *gin.Context) {
    var content content
    if err := ctx.ShouldBindJSON(&content); err != nil {
        ctx.Error(err)
        ctx.AbortWithStatus(http.StatusBadRequest)
        return
    }
    ctx.JSON(http.StatusOK, content)
}