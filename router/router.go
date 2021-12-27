package router

import (
	"j2ee/api"
	"j2ee/middleware"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// store := cookie.NewStore([]byte("secret_string..."))
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	// store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	r.Use(sessions.Sessions("mysession", store))

	r.Use(middleware.Cors())

	// 通用操作
	v0 := r.Group("api/v0")
	{
		// 验证码
		v0.GET("get_captcha", api.GetCaptcha)
		v0.GET("captcha/:captchaID", api.GetCaptchaImage)
	}

	// 用户操作
	v1 := r.Group("/api/v1")
	{
		// 用户注册与登录
		v1.POST("register", api.UserRegister)
		v1.POST("login", api.UserLogin)

		// 登录信息验证
		auth := v1.Group("/")
		auth.Use(middleware.JWT())
		{
			// 验证token（用中间件）
			auth.GET("check_token", api.Ping)

			// 用户信息操作
			auth.GET("info", api.UserInfoGet)

			// 订单操作
			auth.POST("order", api.OrderCreate)
			auth.GET("order", api.OrderGet)
			auth.GET("order/:order_id", api.OrderDetail)
			auth.POST("order_rate", api.OrderRate)

			// 常用地址操作
			auth.POST("address", api.AddressCreate)
			auth.GET("address/:id", api.AddressGet)
			auth.PUT("address", api.AddressUpdate)
			auth.PUT("address_pin", api.AddressUpdatePin)
			auth.DELETE("address", api.AddressDelete)

			// 货车操作
			auth.GET("truck", api.TruckGetByUser)
		}
	}

	v2 := r.Group("/api/v2")
	{
		auth2 := v2.Group("/")
		auth2.Use(middleware.JWTAdmin())
		{
			// 管理员信息操作
			auth2.GET("info", api.AdminInfoGet)

			// 用户操作
			auth2.GET("user", api.AdminUserGet)
			auth2.POST("user/status", api.AdminUserStatus)

			// 订单操作
			auth2.GET("order", api.AdminOrderGet)
			auth2.GET("order/:order_id", api.AdminOrderDetail)

			// 货车操作
			auth2.GET("truck", api.TruckGet)
			auth2.POST("truck/create", api.TruckCreate)
			auth2.POST("truck/allocate", api.TruckAllocate)
		}
	}

	// 404 NotFound
	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"status": 404,
	// 		"error": "404, page not exists!",
	// 	})
	// })

	return r
}
