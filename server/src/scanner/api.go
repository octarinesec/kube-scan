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
	router.GET("refreshing_status", getRefreshingStatus)
	router.POST("refresh", runRefreshState)
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

func getRefreshingStatus(c *gin.Context) {
	type refreshStatus struct {
		LastRefresh int64 `json:"lastRefresh"`
		Refreshing  bool  `json:"refreshing"`
	}
	refreshing, lastRefresh := getRefresh()
	c.JSON(http.StatusOK, refreshStatus{Refreshing: refreshing, LastRefresh: lastRefresh})
}

func runRefreshState(c *gin.Context) {
	go tryRefreshState()
	c.Status(http.StatusNoContent)
}
