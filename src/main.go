package main

import (
	"github.com/gdgtoledo/linneo/src/routes"
	"github.com/gdgtoledo/linneo/src/service"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func handleSearch(c *gin.Context) {
	r, _ := service.Search("my query")
}

func handleDelete(c *gin.Context) {
	id := c.Param("id")
	r, _ := service.Delete(id)
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
