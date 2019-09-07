package main

import (
	"net/http"

	"github.com/gdgtoledo/linneo/dao"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/plants", func(c *gin.Context) {
		res, err := dao.Search("plants", map[string]interface{}{})
		log.WithFields(log.Fields{
			"result": res,
		}).Info("Query Result")
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Error querying database")
		}
		c.String(http.StatusNoContent, "There are no plants in the primary storage")
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
