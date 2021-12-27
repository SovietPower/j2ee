package api

import (
	"j2ee/service"
	"j2ee/util/logging"

	"github.com/gin-gonic/gin"
)

// AddressCreate 新建地址
func AddressCreate(c *gin.Context) {
	service := service.AddressCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddressCreate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// AddressGet 获取地址
func AddressGet(c *gin.Context) {
	service := service.AddressGetService{}
	res := service.AddressGet(c.Param("id"))
	c.JSON(200, res)
}

// AddressUpdate 修改地址
func AddressUpdate(c *gin.Context) {
	service := service.AddressUpdateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddressUpdate()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// AddressUpdatePin 修改地址的置顶与否（因为简单常用所以单独设置）
func AddressUpdatePin(c *gin.Context) {
	service := service.AddressUpdatePinService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddressUpdatePin()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// AddressDelete 删除地址
func AddressDelete(c *gin.Context) {
	service := service.AddressDeleteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddressDelete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
