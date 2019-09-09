package main

import (
	"net/http"

	"github.com/gdgtoledo/linneo/dao"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmlogrus"
)

func init() {
	// apmlogrus.Hook will send "error", "panic", and "fatal"
	// level log messages to Elastic APM.
	log.AddHook(&apmlogrus.Hook{})
}

func searchPlants(c *gin.Context) {
	// apmlogrus.TraceContext extracts the transaction and span (if any) from the given context,
	// and returns logrus.Fields containing the trace, transaction, and span IDs.
	traceContextFields := apmlogrus.TraceContext(c)
	log.WithFields(traceContextFields).Debug("handling request")

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
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))

	r.GET("/plants", searchPlants)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
