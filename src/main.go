package main

import (
	dao "github.com/gdgtoledo/linneo/src/elastic"
	"github.com/gdgtoledo/linneo/src/routes"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func handleSearch(c *gin.Context) {
	r, _ := dao.Search("my query")
}

func handleDelete(c *gin.Context) {
	id := c.Param("id")
	r, _ := dao.Delete(id)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.GET(routes.Plants, handleSearch)
	r.DELETE(routes.Plant, handleDelete)
	return r
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
