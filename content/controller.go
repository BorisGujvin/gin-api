package content

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BorisGujvin/gin-api/error"
	"github.com/gin-gonic/gin"
)

func GetContents(ctx *gin.Context) {
	name, present := ctx.GetQuery("name")
	if !present {
		err := error.NewHttpError("Query parameter not found", "name query parameter is required", http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	response := make(map[string]string, 0)
	response["name"] = name
	var result = GetContentList()
	ctx.JSON(http.StatusOK, result)
}

func PostContents(ctx *gin.Context) {
	var content content
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.Error(error.NewHttpError("Bad request", "body is invalid", http.StatusBadRequest))
		return
	}
	var result = CreateContent(content.Name)
	ctx.JSON(http.StatusOK, result)
}

func ConsumeFile(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		err := error.NewHttpError("No file in form data", "file parameter is required", http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	csvFileToImport, err := fileHeader.Open()
	if err != nil {
		ctx.Error(error.NewHttpError("No file in fileHeader", "file is required", http.StatusBadRequest))
		return
	}
	defer csvFileToImport.Close()

	fileBytes, err := ioutil.ReadAll(csvFileToImport)
	if err != nil {
		ctx.Error(error.NewHttpError("Cannot read file data", "inner error", http.StatusInternalServerError))
		return
	}

	f, err := os.Create("../../storage/" + fileHeader.Filename)
	if err != nil {
		ctx.Error(error.NewHttpError("Cannot create file", "inner error", http.StatusInternalServerError))
	}

	_, err = f.Write(fileBytes)
	if err != nil {
		ctx.Error(error.NewHttpError("Cannot write file data", "inner error", http.StatusInternalServerError))
		return
	}
	f.Close()

	ctx.JSON(http.StatusOK, nil)
}
