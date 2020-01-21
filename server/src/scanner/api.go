package scanner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kube-scan/rest"
	"net/http"
)

func InitApi(port int) error {
	router := gin.Default()

	router.GET("risks", getRisk)
	return router.Run(fmt.Sprintf(":%v", port))
}

func getRisk(c *gin.Context) {
	if ClusterState == nil {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	riskData := rest.GetClusterRiskWorkloads(ClusterState)
	c.JSON(http.StatusOK, rest.ClusterRiskData{Data: riskData})
}
