package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hapiman/ke/controllers"
	"github.com/hapiman/ke/utils"
)

func FetchDailyStats(c *gin.Context) {
	avgStr, yestTxNum := controllers.FetchDailyStats()
	data := struct {
		AvgStr    string `json:"avgStr"`
		YestTxNum string `json:"yestTxNum"`
	}{
		AvgStr:    avgStr,
		YestTxNum: yestTxNum,
	}
	succ := &utils.RespSucc{
		Error:   nil,
		Success: true,
		Data:    data,
	}
	c.JSON(200, succ)
}

func FetchDailyNew(c *gin.Context) {
	houseList := controllers.FetchDailyNew("")
	succ := &utils.RespSucc{
		Error:   nil,
		Success: true,
		Data: map[string]interface{}{
			"list": houseList,
		},
	}
	c.JSON(200, succ)
}

func FetchDailyTxList(c *gin.Context) {
	houseList := controllers.FetchDailyTxList("")
	succ := &utils.RespSucc{
		Error:   nil,
		Success: true,
		Data: map[string]interface{}{
			"list": houseList,
		},
	}
	c.JSON(200, succ)
}

func main() {
	r := gin.Default()
	r.GET("/ke/api/v1/daily/stats", FetchDailyStats)
	r.GET("/ke/api/v1/daily/new", FetchDailyNew)
	r.GET("/ke/api/v1/daily/txlist", FetchDailyTxList)
	r.Run(fmt.Sprintf(":%d", utils.Port))
}
