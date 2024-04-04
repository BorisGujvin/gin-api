package content

import (
    "fmt"
    "net/http"
    "io/ioutil"
//     "os"
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

func ConsumeFile(ctx *gin.Context) {
    fileHeader, err := ctx.FormFile("file")
    if err != nil {
        ctx.Error(err)
        return
    }

    //Open received file
    csvFileToImport, err := fileHeader.Open()
    if err != nil {
        ctx.Error(err)
        return
    }
    defer csvFileToImport.Close()

    //Create temp file
    tempFile, err := ioutil.TempFile("", fileHeader.Filename)
    if err != nil {
        ctx.Error(err)
        return
    }
    defer tempFile.Close()

    //Delete temp file after importing
   // defer os.Remove(tempFile.Name())

    //Write data from received file to temp file
    fileBytes, err := ioutil.ReadAll(csvFileToImport)
    if err != nil {
        ctx.Error(err)
        return
    }
    _, err = tempFile.Write(fileBytes)
    if err != nil {
        ctx.Error(err)
        return
    }

    ctx.JSON(http.StatusOK, string(fileBytes))
}