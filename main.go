package main

import (
	"net/http"

	"github.com/gdgtoledo/linneo/plants"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var routes = map[string]string{"plants": "/plants", "plant": "/plants/:id"}

func handleSearchItems(c *gin.Context) {
	searchQueryByIndexName := plants.SearchQueryByIndexName{
		IndexName: "plants",
		Query:     map[string]interface{}{},
	}

	res, err := plants.Search(searchQueryByIndexName)

	log.WithFields(log.Fields{
		"result": res,
	}).Info("Query Result")

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error querying database")
	}

	c.String(http.StatusNoContent, "There are no plants in the primary storage")
}

func handleDeleteItem(c *gin.Context) {
	var plant plants.Model
	result, err := plants.Delete(plant.ID)

	log.WithFields(log.Fields{
		"result": result,
	}).Info("Delete Query Result")

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error deleting a plant")
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET(routes["plants"], handleSearchItems)
	r.DELETE(routes["plant"], handleDeleteItem)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
