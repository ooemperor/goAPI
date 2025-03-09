package api

import (
	"github.com/gin-gonic/gin"
	"goAPI/src/adapters"
	"net/http"
	"strconv"
)

type IBaseAPI interface {
	GET(ctx *gin.Context) any
	GetAll(ctx *gin.Context) any
}

type BaseApi struct {
	IBaseAPI
	adapter adapters.IBaseAdapter
}

func (b BaseApi) GET(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.ParseUint(idString, 10, 64)
	ctx.IndentedJSON(http.StatusOK, b.adapter.Get(uint(id)))
}

func (b BaseApi) DELETE(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.ParseUint(idString, 10, 64)
	ctx.IndentedJSON(http.StatusNoContent, b.adapter.Delete(uint(id)))
}

func (b BaseApi) DeleteAll(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusNoContent, b.adapter.DeleteAll())
}

func (b BaseApi) GetAll(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, b.adapter.GetAll())
}

func newBaseApi(adapter adapters.IBaseAdapter) BaseApi {
	return BaseApi{adapter: adapter}
}
