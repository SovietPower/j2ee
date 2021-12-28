package api

import (
	"j2ee/util/logging"
	"j2ee/util/sdk"

	"github.com/gin-gonic/gin"
)

func GetDistricts(c *gin.Context) {
	var service sdk.GetDistrictsService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetDistricts()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(200, ErrorResponse(err))
	}
}
